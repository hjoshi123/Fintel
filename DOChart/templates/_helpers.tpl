{{/*
Expand the name of the chart.
*/}}
{{- define "DOChart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "DOChart.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "DOChart.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "DOChart.labels" -}}
helm.sh/chart: {{ include "DOChart.chart" . }}
{{ include "DOChart.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "DOChart.selectorLabels" -}}
app.kubernetes.io/name: {{ include "DOChart.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "DOChart.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "DOChart.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "DOChart.componentname" -}}
{{- $global := index . 0 -}}
{{- $component := index . 1 | trimPrefix "-" -}}
{{- printf "%s-%s" (include "DOChart.fullname" $global | trunc (sub 62 (len $component) | int) | trimSuffix "-" ) $component | trimSuffix "-" -}}
{{- end -}}

{{- define "DOChart.labels.backend.api" -}}
{{- $global := index . 0 -}}
{{- $component := index . 1 | trimPrefix "-" -}}
{{ include "DOChart.labels" $global }}
{{- if .Values.backend.api.componentName }}
app.kubernetes.io/component: {{ printf "%s-%s" .Values.backend.api.componentName $component }}
{{- end }}
{{- end }}


{{- define "DOChart.labels.backend.pubsub" -}}
{{- $global := index . 0 -}}
{{- $component := index . 1 | trimPrefix "-" -}}
{{ include "DOChart.labels" $global }}
{{- if .Values.backend.pubsub.componentName }}
app.kubernetes.io/component: {{ printf "%s-%s" .Values.backend.pubsub.componentName $component }}
{{- end }}
{{- end }}

{{- define "DOChart.labels.frontend" -}}
{{ include "DOChart.labels" . }}
{{- if .Values.frontend.componentName }}
app.kubernetes.io/component: {{ .Values.frontend.componentName }}
{{- end }}
{{- end }}
