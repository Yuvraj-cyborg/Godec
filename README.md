## godec

A tiny image codec in Go: grayscale → delta encoding → run-length encoding (RLE), packed into a custom `.gdc` binary. The decoder runs the same steps in reverse for a full round trip.

<img src="./assets/architecture.png" alt="Codec pipeline" width="800">

Toy project only not a production format.
