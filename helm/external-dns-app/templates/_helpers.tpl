{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "external-dns.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Giantswarm labels
*/}}
{{- define "giantswarm.labels" -}}
giantswarm.io/service-type: "{{ .Values.serviceType }}"
application.giantswarm.io/team: {{ index .Chart.Annotations "application.giantswarm.io/team" | quote }}
{{- end -}}

{{/*
CRD install annotations.
*/}}
{{- define "annotations.crd" -}}
helm.sh/hook: pre-install,pre-upgrade
helm.sh/hook-delete-policy: before-hook-creation,hook-succeeded
{{- end -}}

{{/*
Create the list of domains to update
*/}}
{{- define "domain.list" }}
{{- if not (eq (kindOf .Values.externalDNS.domainFilterList) "invalid") }}
{{- if eq .Values.aws.access "external" }}
{{- printf "- --domain-filter=%s\n" .Values.aws.baseDomain }}
{{/* kindOf nil == "invalid" */}}
{{- else }}
{{- printf "- --domain-filter=%s\n" .Values.baseDomain }}
{{- end }}
{{- range .Values.externalDNS.domainFilterList }}
{{ printf "- --domain-filter=%s\n" . }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Set the zone type when running on AWS
*/}}
{{- define "zone.type" -}}
{{- if eq .Values.aws.access "external" }}
{{ printf "- --aws-zone-type=%s\n" "public" }}
{{- else }}
{{- if .Values.aws.zoneType }}
{{ printf "- --aws-zone-type=%s\n" .Values.aws.zoneType }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Set the role name for KIAM
*/}}
{{- define "aws.iam.role" -}}
{{- if .Values.aws.iam.customRoleName }}
{{- printf "%s" .Values.aws.iam.customRoleName }}
{{- else }}
{{- if eq .Values.aws.access "internal" }}
{{- printf "%s-Route53Manager-Role" .Values.clusterID }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Set the annotation filter.
*/}}
{{- define "annotation.filter" -}}
{{- if .Values.NetExporter -}}
{{/* if this value is present then the app was installed
from the default catalog and is therefore a default app */}}
{{- print "giantswarm.io/external-dns=managed" }}
{{- else -}}
{{/* the customer must provide their own */}}
{{- printf "%s" .Values.externalDNS.annotationFilter }}
{{- end -}}
{{- end -}}

{{/*
Set the txt owner ID.
*/}}
{{- define "txt.owner.id" -}}
{{- if .Values.NetExporter -}}
{{/* if this value is present then the app was installed
from the default catalog and is therefore a default app */}}
{{- print "giantswarm-io-external-dns" }}
{{- else -}}
{{/* the customer must provide their own */}}
{{- printf "%s" .Values.externalDNS.registry.txtOwnerID }}
{{- end -}}
{{- end -}}

{{/*
Set the txt record prefix.
*/}}
{{- define "txt.prefix" -}}
{{- if .Values.NetExporter -}}
{{/* if this value is present then the app was installed
from the default catalog and is therefore a default app */}}
{{- printf "%s" .Values.clusterID }}
{{- else -}}
{{/* the customer must provide their own */}}
{{- printf "%s" .Values.externalDNS.registry.txtPrefix }}
{{- end -}}
{{- end -}}

{{/*
Set provider specific flags (VMWare clusters use Route53 for DNS). Ternary
function returns `aws` if provider is vmware; otherwise it returns
the value of .Values.provider.
*/}}
{{- define "dnsProvider.flags" -}}
{{- $dnsProvider := .Values.provider -}}
{{- if eq .Values.provider "vmware" }}
{{- $dnsProvider = "aws" -}}
{{- else if eq .Values.provider "capa" }}
{{- $dnsProvider = "aws" -}}
{{- else if eq .Values.provider "gcp" }}
{{- $dnsProvider = "google" -}}
{{- end }}
{{- printf "- --provider=%s" $dnsProvider }}

{{- if eq $dnsProvider "aws" }}
{{ include "zone.type" . }}
{{- if or .Values.aws.preferCNAME (eq .Values.aws.access "external") }}
- --aws-prefer-cname
{{- end }}
{{- if .Values.aws.batchChangeSize }}
- --aws-batch-change-size={{ .Values.aws.batchChangeSize }}
{{- end }}
{{- if .Values.aws.batchChangeInterval }}
- --aws-batch-change-interval={{ .Values.aws.batchChangeInterval }}
{{- end }}
{{ include "domain.list" . }}
{{- end }}

{{- if eq .Values.provider "vmware" }}
- --source=crd
{{- end }}

{{- if eq .Values.provider "azure" }}
- --azure-config-file=/config/azure.yaml
{{- end }}

{{- end }}

{{/*
Validate certain values and fail if they are incorrect
*/}}

{{/* Compile failure message. This function is called
in NOTES.txt */}}
{{- define "externalDNS.validateValues" -}}
{{- $messages := list -}}
{{- $messages := append $messages (include "validateValues.provider" .) -}}
{{- $messages := append $messages (include "validateValues.zoneType" .) -}}
{{- $messages := without $messages "" -}}
{{- $message := join "\n" $messages -}}

{{/* Print failure message(s) */}}
{{- if $message -}}
{{-   printf "\n\nVALUES VALIDATION:\n\n%s" $message | fail -}}
{{- end -}}
{{- end -}}

{{/*
Validate that the provider makes sense

'inmemory' is only used when running integration tests, so we
don't expose that value to the user in the error message.
*/}}
{{- define "validateValues.provider" -}}
{{- if and (ne .Values.provider "aws") (ne .Values.provider "azure") (ne .Values.provider "capa") (ne .Values.provider "gcp") (ne .Values.provider "vmware") (ne .Values.provider "inmemory") -}}
external-dns: provider
    Incorrect value provided. Valid values are either 'aws', 'azure', 'capa', 'gcp' or 'vmware'.
{{- end -}}
{{- end -}}

{{/* Ensure hosted zone type makes sense */}}
{{- define "validateValues.zoneType" -}}
{{- if .Values.aws.zoneType -}}
{{- if and (ne .Values.aws.zoneType "public") (ne .Values.aws.zoneType "private") -}}
external-dns: aws.zoneType
    AWS hosted zone type must be either 'public' or 'private'.
{{- end -}}
{{- end -}}
{{- end -}}


{{/*
Set Giant Swarm podAnnotations.
*/}}
{{- define "giantswarm.podAnnotations" -}}
{{- if and (or (eq .Values.provider "aws") (eq .Values.provider "capa")) (eq .Values.aws.access "internal") ( eq .Values.aws.irsa "false") }}
{{- $_ := set .Values.podAnnotations "iam.amazonaws.com/role" (tpl "{{ template \"aws.iam.role\" . }}" .) }}
{{- end }}
{{- end -}}

{{/*
Set Giant Swarm serviceAccountAnnotations.
*/}}
{{- define "giantswarm.serviceAccountAnnotations" -}}
{{- if and (eq .Values.provider "aws") (eq .Values.aws.irsa "true") (eq .Values.aws.access "internal") (not (hasKey .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn")) }}
{{- $_ := set .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn" (tpl "arn:aws:iam::{{ .Values.aws.accountID }}:role/{{ template \"aws.iam.role\" . }}" .) }}
{{- else if and (eq .Values.provider "capa") (eq .Values.aws.access "internal") (not (hasKey .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn")) }}
{{- $_ := set .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn" (include "aws.iam.role" .) }}
{{- else if and (eq .Values.provider "gcp") (.Values.gcpProject) (not (hasKey .Values.serviceAccount.annotations "giantswarm.io/gcp-service-account")) }}
{{- $_ := set .Values.serviceAccount.annotations "giantswarm.io/gcp-service-account" (tpl "external-dns-app@{{ .Values.gcpProject }}.iam.gserviceaccount.com" .) }}
{{- end }}
{{- end -}}

{{/*
Set Giant Swarm env for Deployment.
*/}}
{{- define "giantswarm.deploymentEnv" -}}
{{- if and .Values.externalDNS.aws_access_key_id .Values.externalDNS.aws_secret_access_key }}
- name: AWS_ACCESS_KEY_ID
  valueFrom:
  secretKeyRef:
    name: {{ .Release.Name }}-route53-credentials
    key: aws_access_key_id
- name: AWS_SECRET_ACCESS_KEY
  valueFrom:
  secretKeyRef:
    name: {{ .Release.Name }}-route53-credentials
    key: aws_secret_access_key
{{- end }}
{{- with .Values.aws.region }}
- name: AWS_DEFAULT_REGION
  value: {{ . }}
{{- end }}
{{- $proxy := deepCopy .Values.cluster.proxy |  mustMerge .Values.proxy -}}
{{- if and $proxy.noProxy $proxy.http $proxy.https }}
- name: NO_PROXY
  value: {{ $proxy.noProxy }}
- name: no_proxy
  value: {{ $proxy.noProxy }}
- name: HTTP_PROXY
  value: {{ $proxy.http }}
- name: http_proxy
  value: {{ $proxy.http }}
- name: HTTPS_PROXY
  value: {{ $proxy.https }}
- name: https_proxy
  value: {{ $proxy.https }}
{{- end }}
{{- end -}}

{{/*
Set Giant Swarm initContainers.
*/}}
{{- define "giantswarm.deploymentInitContainers" -}}
{{- if and (or (eq .Values.provider "aws") (eq .Values.provider "capa")) (eq .Values.aws.access "internal") ( eq .Values.aws.irsa "false") }}
- name: wait-for-iam-role
  image: {{ .Values.global.image.registry }}/giantswarm/alpine:3.16.2
  command:
    - /bin/sh
    - -c
    - counter=5; while ! wget -qO- http://169.254.169.254/latest/meta-data/iam/security-credentials/ | grep {{ template "aws.iam.role" . }}; do echo 'Waiting for iam-role to be available...'; sleep 5; let "counter-=1"  ; if [ "$counter" -eq "0" ]; then exit 1; fi; done
{{- end }}
{{- if eq .Values.provider "azure" }}
- name: copy-azure-config-file
  image: {{ .Values.global.image.registry }}/giantswarm/alpine:3.16.2-python3
  command:
    - /bin/sh
    - -c
    # GS clusters have the cloud config file in /etc/kubernetes/config/azure.yaml and we can use it as-is so we just copy it to the desired position.
    # CAPZ clusters use a JSON file so we convert it to yaml and save it to the desired position.
    - if [ -f /etc/kubernetes/config/azure.yaml ]; then
      cp /etc/kubernetes/config/azure.yaml /config/azure.yaml;
      else
      cat /etc/kubernetes/azure.json | python3 -c 'import sys, yaml, json; print(yaml.dump(json.loads(sys.stdin.read())))' > /config/azure.yaml;
      fi
  volumeMounts:
    - mountPath: /etc/kubernetes
      name: etc-kubernetes
      readOnly: true
    - mountPath: /config
      name: config
{{- end }}
{{- end -}}

{{/*
Upstream chart helpers.
*/}}

{{/*
Expand the name of the chart.
*/}}
{{- define "external-dns.name" -}}
{{- default "external-dns" .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "external-dns.labels" -}}
helm.sh/chart: {{ include "external-dns.chart" . }}
app.kubernetes.io/name: {{ include "external-dns.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{ include "external-dns.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- with .Values.commonLabels }}
{{ toYaml . }}
{{- end }}
{{ include "giantswarm.labels" . }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "external-dns.selectorLabels" -}}
app: {{ .Release.Name | quote }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
Unless there is a override, we use the release name as the full name.
*/}}
{{- define "external-dns.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "external-dns.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "external-dns.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
