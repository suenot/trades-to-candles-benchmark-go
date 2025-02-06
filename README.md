# Trades to Candles Benchmark

A comprehensive benchmark and comparison of different Go libraries for trade data aggregation and candlestick conversion.

## Overview

This project aims to evaluate and compare various Go libraries that handle:
- Trade data to candlestick conversion
- Candlestick timeframe aggregation (e.g., 1m to 5m)

## TL;DR
```
go run cmd/benchmark/main.go -data data/BTCUSDT-trades-2025-01-30.csv
Loaded 3990075 trades
Time range: from 2025-01-30 03:00:00.232125 +0300 MSK to 2025-01-31 02:59:59.778505 +0300 MSK
Total duration: 23h59m59.54638s

Candle Analysis:
Number of candles: 1440
First candle: Candle{Time: 2025-01-31 01:11:00 +0300 MSK, O: 105052.01, H: 105099.99, L: 105052, C: 105099.98, V: 13.422909999999952}
Last candle: Candle{Time: 2025-01-30 19:46:00 +0300 MSK, O: 105768.68, H: 105768.69, L: 105684.8, C: 105730.21, V: 14.005579999999599}
Warning: Expected 1439 candles, but got 1440

First 5 candles:
Candle{Time: 2025-01-31 01:11:00 +0300 MSK, O: 105052.01, H: 105099.99, L: 105052, C: 105099.98, V: 13.422909999999952}
Candle{Time: 2025-01-30 04:29:00 +0300 MSK, O: 104432.91, H: 104464, L: 104372.22, C: 104372.22, V: 38.442009999999186}
Candle{Time: 2025-01-30 10:51:00 +0300 MSK, O: 104939.9, H: 104969.51, L: 104934.4, C: 104969.51, V: 2.6940099999999707}
Candle{Time: 2025-01-30 11:18:00 +0300 MSK, O: 105000, H: 105016.8, L: 104989.99, C: 105016.8, V: 13.868549999999917}
Candle{Time: 2025-01-30 14:01:00 +0300 MSK, O: 105267.6, H: 105283, L: 105247.35, C: 105283, V: 3.785679999999966}

Last 5 candles:
Candle{Time: 2025-01-30 08:25:00 +0300 MSK, O: 105245.29, H: 105300, L: 105245.28, C: 105299.99, V: 2.931809999999909}
Candle{Time: 2025-01-30 08:58:00 +0300 MSK, O: 105282.95, H: 105323, L: 105273.14, C: 105322.99, V: 14.515089999999976}
Candle{Time: 2025-01-30 11:41:00 +0300 MSK, O: 105048.2, H: 105090.78, L: 105048.19, C: 105090.78, V: 2.589949999999987}
Candle{Time: 2025-01-30 13:28:00 +0300 MSK, O: 105250.38, H: 105258.69, L: 105250.37, C: 105258.01, V: 3.71188999999999}
Candle{Time: 2025-01-30 19:46:00 +0300 MSK, O: 105768.68, H: 105768.69, L: 105684.8, C: 105730.21, V: 14.005579999999599}

Benchmark Results Comparison:
=============================

Library: minute_aggregator
Execution Time: 378.542125ms
Memory Usage: 0 MB
Max Memory Usage: 18 MB
Processed Trades: 3990075
Candles Created: 1440
```
## Libraries Under Test

### Trades to Candles
- [techan](https://github.com/sdcoffey/techan) - Technical analysis library
- [marketstore](https://github.com/alpacahq/marketstore) - Financial data timestore
- [go-trader](https://github.com/saniales/golang-crypto-trading-bot) - Trading bot with OHLCV support
- [golangcandles](https://github.com/complimenti/golangcandles) - Lightweight candle processing
- [go_trade_aggregation](https://github.com/MathisWellmann/go_trade_aggregation) - Specialized trade aggregation
- [cinar/indicator](https://github.com/cinar/indicator) - Simple and efficient
- [go-exchanges](https://github.com/go-numb/go-exchanges) - Exchange integration focused

### Candles to Candles
- [timeframe-aggregation-rs](https://github.com/suenot/timeframe-aggregation-rs) - Rust library with Go bindings
- [ta](https://github.com/miaolz123/ta) - Technical analysis with timeframe conversion

## Benchmark Details

### Test Dataset
- Instrument: BTCUSDT
- Date: 2025-01-30
- Format: ZIP archive containing trade data

### Test Scenarios
1. Trades to Candles:
   - Aggregation to 1m candles
   - Aggregation to 5m candles
   - Large dataset processing (>1M trades)

2. Candles to Candles:
   - 1m -> 5m conversion
   - 1m -> 1h conversion
   - Edge cases testing

### Metrics
- Execution time
- CPU usage
- Memory consumption
- Result accuracy

## Project Structure 

## Current Status
- [x] Basic infrastructure setup
- [x] Initial library integration
- [x] Benchmark implementation
- [x] First benchmark results

## Benchmark Results

### go_trade_aggregation
First benchmark results:
- Test Environment:
  * Hardware: Apple MacBook Air M1 (2020)
  * RAM: 16GB
  * OS: macOS
- Dataset:
  * Symbol: BTCUSDT
  * Size: 1.2M trades
  * Period: 24 hours
- Characteristics:
  * Stable performance on large datasets
  * Efficient memory usage
  * Clean API integration

### Other Libraries
Benchmarks for other libraries are in progress:
- [ ] golangcandles
- [ ] techan
- [ ] marketstore
- [ ] cinar/indicator
- [ ] go-trader
- [ ] go-exchanges
- [ ] ta