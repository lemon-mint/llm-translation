go generate && \
  sleep 0.1 && \
  cd web && \
  bun run build && \
  cd .. && \
  go build -o ./tmp/main
