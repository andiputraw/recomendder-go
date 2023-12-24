# Recomendder go (not completed)

Recomenddation system from scratch

# Dataset

1. [Geeks for geek article](https://www.kaggle.com/datasets/naidukarthi2193/geeks-for-geeks-articles-dataset). 50 Mb dataset. this is the recomendded set
2. [190K+ medium article](https://www.kaggle.com/datasets/fabiochiusano/medium-articles?resource=download). 1GB . if you are using this. make sure to uncomment this line `docs := tf.DocFromCsvMultithread(reader, 100)` as it will be very slow to parse. after that change constant value in `tf.go`

# Compiling

Make you have golang and C compiler

I am using [Zig](https://ziglang.org/) for C compiler. if you are using different complier. change the `CC` variable in `run.sh` files

after that. run `run.sh` in your terminal

```sh
./run.sh
```

## References

- [TF-IDF from wikipedia](https://en.wikipedia.org/wiki/Tf%E2%80%93idf)
- [Search engine in Rust from Tsoding (Youtube)](https://youtu.be/hm5xOJiVEeg?si=NQOA5c0W8m2NDx4-)
- [Medium article (I only read the title)](https://medium.com/geekculture/understanding-tf-idf-and-cosine-similarity-for-recommendation-engine-64d8b51aa9f9)
- [Cosine similarity explanation By StatQuest (Youtube)](https://youtu.be/e9U0QAFbfLI?si=b3ytcZQs7KpHh4tw)
