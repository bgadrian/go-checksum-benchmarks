# Checksums benchmarks
This project was used for research, I want to know which checksum to use (best for all):
 * small text
 * small binary
 * encrypted data
 
 What are we looking for:
  * very fast
  * low collision rate for small messages
  
  We do not care about the cryptographic hashes, I only need an unique hash for each input.
  
  ## Algorithms:
   * CRC64 checksum https://golang.org/pkg/hash/crc64/
   * FNV1 64&128 https://golang.org/pkg/hash/fnv/
   * Murmur3 sum 64&128 github.com/spaolacci/murmur3
   * Jenkins 64&128- SpookyHash https://github.com/dgryski/go-spooky
   Jenkins is a family of fast hash algorithms https://en.m.wikipedia.org/wiki/Jenkins_hash_function
   
   