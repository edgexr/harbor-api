# harbor-api
Generated go models for goharbor/harbor

Generated from:
```bash
git clone https://github.com/goharbor/harbor
cd harbor
# Currently, release is v2.6.0
git checkout <release>
make gen_apis
cp -R src/server/v2.0/models ../../edgexr/harbor-api
```
