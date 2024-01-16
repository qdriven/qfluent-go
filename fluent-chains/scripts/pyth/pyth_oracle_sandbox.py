"""
```bash
export SOLANA_RPC_PORT=8899
export PYTHD_WS_PORT=8910

docker build . -t oracle-sandbox --build-arg SOLANA_RPC_PORT=$SOLANA_RPC_PORT --build-arg PYTHD_WS_PORT=$PYTHD_WS_PORT
docker run -p $SOLANA_RPC_PORT:$SOLANA_RPC_PORT -p $PYTHD_WS_PORT:$PYTHD_WS_PORT oracle-sandbox
```
"""