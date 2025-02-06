# Trades to Candles Benchmark

A comprehensive benchmark and comparison of different Go libraries for trade data aggregation and candlestick conversion.

## Overview

This project aims to evaluate and compare various Go libraries that handle:
- Trade data to candlestick conversion
- Candlestick timeframe aggregation (e.g., 1m to 5m)

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