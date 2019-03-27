pkitree
-------

A tool to view relations between certificates in some directory.

Example:

```text
.
├── fa904 [CA] ca.localhost -> o/l/c/ca.localhost-cert.pem
│   ├── d80cd Admin@localhost -> o/l/m/a/Admin@localhost-cert.pem
│   ├── d80cd Admin@localhost -> o/l/o/o/m/a/Admin@localhost-cert.pem
│   ├── 11bf2 orderer.localhost -> o/l/o/o/m/s/orderer.localhost-cert.pem
│   ├── d80cd Admin@localhost -> o/l/u/A/m/a/Admin@localhost-cert.pem
│   └── d80cd Admin@localhost -> o/l/u/A/m/s/Admin@localhost-cert.pem
├── fa904 [CA] ca.localhost -> o/l/m/c/ca.localhost-cert.pem
├── 0c0c6 [CA] tlsca.localhost -> o/l/m/t/tlsca.localhost-cert.pem
│   ├── 77164 orderer.localhost -> o/l/o/o/t/server.crt
│   └── e6ab9 Admin@localhost -> o/l/u/A/t/client.crt
├── fa904 [CA] ca.localhost -> o/l/o/o/m/c/ca.localhost-cert.pem
├── 0c0c6 [CA] tlsca.localhost -> o/l/o/o/m/t/tlsca.localhost-cert.pem
├── 0c0c6 [CA] tlsca.localhost -> o/l/o/o/t/ca.crt
├── 0c0c6 [CA] tlsca.localhost -> o/l/t/tlsca.localhost-cert.pem
├── fa904 [CA] ca.localhost -> o/l/u/A/m/c/ca.localhost-cert.pem
├── 0c0c6 [CA] tlsca.localhost -> o/l/u/A/m/t/tlsca.localhost-cert.pem
├── 0c0c6 [CA] tlsca.localhost -> o/l/u/A/t/ca.crt
├── 1f308 [CA] ca.default-org.localhost -> p/d/c/ca.default-org.localhost-cert.pem
│   ├── 6dd9d Admin@default-org.localhost -> p/d/m/a/Admin@default-org.localhost-cert.pem
│   ├── 6dd9d Admin@default-org.localhost -> p/d/p/p/m/a/Admin@default-org.localhost-cert.pem
│   ├── 4c8ed peer0.default-org.localhost -> p/d/p/p/m/s/peer0.default-org.localhost-cert.pem
│   ├── 6dd9d Admin@default-org.localhost -> p/d/u/A/m/a/Admin@default-org.localhost-cert.pem
│   ├── 6dd9d Admin@default-org.localhost -> p/d/u/A/m/s/Admin@default-org.localhost-cert.pem
│   ├── 55fe5 User1@default-org.localhost -> p/d/u/U/m/a/User1@default-org.localhost-cert.pem
│   └── 55fe5 User1@default-org.localhost -> p/d/u/U/m/s/User1@default-org.localhost-cert.pem
├── 1f308 [CA] ca.default-org.localhost -> p/d/m/c/ca.default-org.localhost-cert.pem
├── 330c5 [CA] tlsca.default-org.localhost -> p/d/m/t/tlsca.default-org.localhost-cert.pem
│   ├── c952b peer0.default-org.localhost -> p/d/p/p/t/server.crt
│   ├── 0cfa2 Admin@default-org.localhost -> p/d/u/A/t/client.crt
│   └── 022a9 User1@default-org.localhost -> p/d/u/U/t/client.crt
├── 1f308 [CA] ca.default-org.localhost -> p/d/p/p/m/c/ca.default-org.localhost-cert.pem
├── 330c5 [CA] tlsca.default-org.localhost -> p/d/p/p/m/t/tlsca.default-org.localhost-cert.pem
├── 330c5 [CA] tlsca.default-org.localhost -> p/d/p/p/t/ca.crt
├── 330c5 [CA] tlsca.default-org.localhost -> p/d/t/tlsca.default-org.localhost-cert.pem
├── 1f308 [CA] ca.default-org.localhost -> p/d/u/A/m/c/ca.default-org.localhost-cert.pem
├── 330c5 [CA] tlsca.default-org.localhost -> p/d/u/A/m/t/tlsca.default-org.localhost-cert.pem
├── 330c5 [CA] tlsca.default-org.localhost -> p/d/u/A/t/ca.crt
├── 1f308 [CA] ca.default-org.localhost -> p/d/u/U/m/c/ca.default-org.localhost-cert.pem
├── 330c5 [CA] tlsca.default-org.localhost -> p/d/u/U/m/t/tlsca.default-org.localhost-cert.pem
└── 330c5 [CA] tlsca.default-org.localhost -> p/d/u/U/t/ca.crt
```
