# gobert
![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)
[![GoDoc](https://godoc.org/github.com/buckhx/gobert?status.svg)](https://godoc.org/github.com/buckhx/gobert)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Under Active Development

Go bindings for operationalizing BERT models. Train in Python, run in Go.

The following advice from Go TensorFlow applies:
```
TensorFlow provides a Go API— particularly useful for loading models created with Python and running them within a Go application.
...
TensorFlow provides APIs for use in Go programs. These APIs are particularly well-suited to loading models created in Python and executing them within a Go application.
...
Be a real gopher, keep it simple! Use Python to define & train models; you can always load trained models and using them with Go later!
```

### Notes

* SeqLen has large impact on performance
* Perf is not great, need to determine if it's from python model or go runtime
* goimports isnt running, so may have messy imports
* model package is WIP

## Examples

* SemanticSearch: Simple search engine from CSV data
* Classifier: exposing model from run_classifier
* Embeddings: returning sentence embeddings
* Raw: Using only tokenize/vocab package

## Packages

### Tokenize

The tokenize package includes methods to create BERT input features. It is fairly stable and can be used independently of the model package.

### Vocab

The vocab package is a simple container for BERT vocabs

###  Model

The model package is an experimental package to work with models exported 

### Python

The python dir includes utilities to export BERT models that can be exposed to the GO runtime.
There is a loose coupling with the model package and exported models interop with the model package.


# TODOs
- [X] Python Embedding
- [X] Python Classifier
- [ ] Go Classifier
- [X] Semantic Search Example
- [ ] Raw Model Example
- [ ] Token Lookup
- [ ] Model Download
- [ ] Documentation
- [ ] Cleanup makefile
- [ ] Test Coverage 
- [ ] Benchmark
- [ ] Binary CMD

# TBD
- [ ] gonum interop
- [ ] first class wrapper API ([][][]float32 -> []Embedding)
- [ ] proto interops
- [ ] pooling strategies
- [ ] batching
- [ ] other tuned models (SentencePrediction, SQUAD)

Current line of thought is to use core lib for raw types and supply a utility API
