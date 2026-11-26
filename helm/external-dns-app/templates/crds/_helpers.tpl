{{- define "external-dns.crdInstaller.name" -}}
{{- printf "%s-crd-installer" (include "external-dns.fullname" . | trunc 48 | trimSuffix "-") }}
{{- end }}

{{- define "external-dns.crdInstaller.selectorLabels" -}}
app.kubernetes.io/name: {{ include "external-dns.crdInstaller.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "external-dns.crdInstaller.labels" -}}
helm.sh/chart: {{ include "external-dns.chart" . }}
{{ include "external-dns.crdInstaller.selectorLabels" . }}
app.kubernetes.io/component: crd-installer
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{ include "giantswarm.labels" . }}
{{- end }}

{{- define "external-dns.crdInstaller.annotations" -}}
"helm.sh/hook": "pre-install,pre-upgrade"
"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded,hook-failed"
{{- end }}
