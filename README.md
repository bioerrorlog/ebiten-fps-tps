# Ebitengine FPS and TPS

Show FPS and TPS with Update()/Draw() in Ebitengine.


## Results

### TPS 60, Display refresh rate 60

```log
2024/11/25 14:41:54 TPS: 59.92, FPS: 59.92
2024/11/25 14:41:54 Update() was called in this sec: 60 times
2024/11/25 14:41:54 Draw() was called in this sec: 60 times

2024/11/25 14:41:55 TPS: 60.09, FPS: 60.09
2024/11/25 14:41:55 Update() was called in this sec: 61 times
2024/11/25 14:41:55 Draw() was called in this sec: 61 times

2024/11/25 14:41:56 TPS: 59.98, FPS: 59.98
2024/11/25 14:41:56 Update() was called in this sec: 61 times
2024/11/25 14:41:56 Draw() was called in this sec: 61 times
```

### TPS 100, Display refresh rate 60

```log
2024/11/25 14:46:47 TPS: 100.32, FPS: 60.00
2024/11/25 14:46:47 Update() was called in this sec: 100 times
2024/11/25 14:46:47 Draw() was called in this sec: 60 times

2024/11/25 14:46:48 TPS: 100.41, FPS: 60.05
2024/11/25 14:46:48 Update() was called in this sec: 100 times
2024/11/25 14:46:48 Draw() was called in this sec: 60 times

2024/11/25 14:46:49 TPS: 99.92, FPS: 59.95
2024/11/25 14:46:49 Update() was called in this sec: 100 times
2024/11/25 14:46:49 Draw() was called in this sec: 60 times
```

### TPS 60, Display refresh rate 50

```log
2024/11/25 15:04:59 TPS: 60.00, FPS: 50.00
2024/11/25 15:04:59 Update() was called in this sec: 60 times
2024/11/25 15:04:59 Draw() was called in this sec: 50 times

2024/11/25 15:05:00 TPS: 59.86, FPS: 50.05
2024/11/25 15:05:00 Update() was called in this sec: 61 times
2024/11/25 15:05:00 Draw() was called in this sec: 51 times

2024/11/25 15:05:01 TPS: 59.94, FPS: 49.95
2024/11/25 15:05:01 Update() was called in this sec: 60 times
2024/11/25 15:05:01 Draw() was called in this sec: 50 times
```

## Reference

- [EbitengineにおけるFPSとTPSの違いを理解する - BioErrorLog Tech Blog](https://www.bioerrorlog.work/entry/ebitengine-fps-tps)
