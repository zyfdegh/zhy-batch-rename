# zhy-batch-rename

zhy has a need to rename directories under `/root/_alllayers` like following

> Before

```bash
.
├── L18
│   ├── R0001D1D5
│   ├── R0001D1D6
│   └── R0001D1D7
└── L19
    ├── R0001D1D1
    ├── R0001D1D2
    └── R0001D1D3
```

> After

```bash
.
├── L18
│   ├── R0001d1d5
│   ├── R0001d1d6
│   └── R0001d1d7
└── L19
    ├── R0001d1d1
    ├── R0001d1d2
    └── R0001d1d3
```

# Build & Run
```sh
go build && sudo ./zhy-batch-rename
```

/root/_alllayers/L18/R0001D1D5 -> /root/_alllayers/L18/R0001d1d5
/root/_alllayers/L18/R0001D1D6 -> /root/_alllayers/L18/R0001d1d6
/root/_alllayers/L18/R0001D1D7 -> /root/_alllayers/L18/R0001d1d7
/root/_alllayers/L19/R0001D1D1 -> /root/_alllayers/L19/R0001d1d1
/root/_alllayers/L19/R0001D1D2 -> /root/_alllayers/L19/R0001d1d2
/root/_alllayers/L19/R0001D1D3 -> /root/_alllayers/L19/R0001d1d3

