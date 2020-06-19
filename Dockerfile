FROM scratch
ENV GOROOT /usr/local
## docker set env APP_MODE = PRODUCTION
ENV APP_MODE production
# docker set env to +7:00 time zone
ENV TZ=Asia/Bangkok

ADD https://golang.org/lib/time/zoneinfo.zip /usr/local/lib/time/
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/ca-certificates.crt
#ADD secret/mch_privkey.pem /secret/mch_privkey.pem
#ADD secret/service_account.json /secret/service_account.json

COPY application /
#COPY admin/html /admin/html
#COPY admin/templates /admin/templates
# COPY admin/assets /admin/assets
#COPY auth/templates /auth/templates
#COPY member/templates /member/templates
#COPY static /static
# COPY auth/rbac_domain_model.conf /auth/rbac_domain_model.conf
# COPY auth/policy_domain.csv /auth/policy_domain.csv

## Add config file with Global variable such as : time server
#ADD config.yml /

EXPOSE 8080

ENTRYPOINT ["/app"]
