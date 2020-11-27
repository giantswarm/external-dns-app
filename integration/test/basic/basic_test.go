// +build k8srequired

package basic

import (
	"context"
	"testing"
	"time"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestBasic ensures that there is a ready external-dns deployment.
func TestBasic(t *testing.T) {
	ctx := context.Background()
	var err error

	// Check external-dns deployment is ready.
	err = checkReadyDeployment(ctx)
	if err != nil {
		t.Fatalf("could not get external-dns: %v", err)
	}
}

func checkReadyDeployment(ctx context.Context) error {
	var err error

	l.LogCtx(ctx, "level", "debug", "message", "waiting for ready deployment")

	o := func() error {
		deploy, err := appTest.K8sClient().AppsV1().Deployments(metav1.NamespaceSystem).Get(ctx, app, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			return microerror.Maskf(executionFailedError, "deployment %#q in %#q not found", app, metav1.NamespaceSystem)
		} else if err != nil {
			return microerror.Mask(err)
		}

		if deploy.Status.ReadyReplicas != *deploy.Spec.Replicas {
			return microerror.Maskf(executionFailedError, "deployment %#q want %d replicas %d ready", deploy.Name, *deploy.Spec.Replicas, deploy.Status.ReadyReplicas)
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "deployment is ready")

	return nil
}
