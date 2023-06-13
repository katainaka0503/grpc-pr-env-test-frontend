# syntax = docker/dockerfile:1.3-labs
FROM nginx:1.24.0-alpine

RUN <<EOF
  curl -L https://github.com/DataDog/nginx-datadog/releases/download/v1.0.0/nginx_1.24.0-alpine-ngx_http_datadog_module.so.tgz -O .
  tar -xvf nginx_1.24.0-alpine-ngx_http_datadog_module.so.tgz -C /usr/lib/nginx/modules
  rm nginx_1.24.0-alpine-ngx_http_datadog_module.so.tgz
  chmod +x /usr/lib/nginx/modules/ngx_http_datadog_module.so
EOF

COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf