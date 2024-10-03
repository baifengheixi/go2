KAFKA=kafka
KAFKA_UI=kafka-ui
KEYCLOAK=keycloak
LOKI_STACK=loki-stack
VAULT=vault

APP_NAMESPACE=app
KAFKA_NAMESPACE=kafka
CONFIG_NAMESPACE=config
MONITOR_NAMESPACE=monitor

.PHONY: ${KAFKA} ${KAFKA_UI} ${KEYCLOAK} ${LOKI_STACK} ${VAULT}

all: ${KAFKA} ${KAFKA_UI} ${VAULT} ${KEYCLOAK} ${LOKI_STACK}

${KAFKA}:
	helm upgrade -i --create-namespace -n ${KAFKA_NAMESPACE} ${KAFKA} ./${KAFKA} -f ./values/${KAFKA}-values.yaml

${KAFKA_UI}:
	helm upgrade -i --create-namespace -n ${KAFKA_NAMESPACE} ${KAFKA_UI} ./${KAFKA_UI} -f ./values/${KAFKA_UI}-values.yaml

${KEYCLOAK}:
	helm upgrade -i --create-namespace -n ${APP_NAMESPACE} ${KEYCLOAK} ./${KEYCLOAK} -f ./values/${KEYCLOAK}-values.yaml

${LOKI_STACK}:
	helm upgrade -i --create-namespace -n ${MONITOR_NAMESPACE} ${LOKI_STACK} ./${LOKI_STACK} -f ./values/${LOKI_STACK}-values.yaml

${VAULT}:
	helm upgrade -i --create-namespace -n ${CONFIG_NAMESPACE} ${VAULT} ./${VAULT} -f ./values/${VAULT}-values.yaml