FROM centos:8.1.1911

WORKDIR /MuxiFresh-Be-2.0/gateway

COPY gateway gateway
COPY etc/gateway.yaml etc/gateway.yaml

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]