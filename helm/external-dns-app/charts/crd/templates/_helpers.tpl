{{/*
CRD install annotations.
*/}}
{{- define "crd.annotations" -}}
helm.sh/hook: pre-install,pre-upgrade
helm.sh/hook-delete-policy: before-hook-creation,hook-succeeded
{{- end -}}

{{/*
Common labels
*/}}
{{- define "crd.labels" -}}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
giantswarm.io/service-type: "{{ .Values.serviceType }}"
application.giantswarm.io/team: {{ index .Chart.Annotations "application.giantswarm.io/team" | quote }}
{{ include "crd.selectorLabels" . }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "crd.selectorLabels" -}}
app.kubernetes.io/name: external-dns-crd-install
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/component: crd-install
{{- end }}
