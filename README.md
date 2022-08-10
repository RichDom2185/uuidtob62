# uuidtob62

`uuidtob62` is a command-line utility to convert a UUID in standard (hexadecimal) representation to a base-62 string. It prints the output base-62 string to standard output. `uuidtob62` is written in Golang.

The main purpose of this utility is to provide a more compact representation of a UUID in use cases where there may be limited space available. A normal UUID consists of a total of 36 characters: 32 hexadecimal characters and 4 hyphens. The positions of these 4 hyphens are fixed, hence they can be dropped. The remaining characters are parsed as a base-16 number and converted to base 62 with the following radix, as specified in the [big package documentation](https://pkg.go.dev/math/big?utm_source=gopls#Int.Text):

```
0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
```

Converting the UUID to base 62 results in a 22-character string, saving 39% space. In small, low-resolutions screens, this can make all the difference e.g. when trying to display a unique QR code.

## Usage

- Read UUID from a file:

  ```bash
  $ uuidtob62 <file_name>
  ```

- Pass UUID string to standard input:

  ```bash
  $ echo '123e4567-e89b-12d3-a456-426614174000' > uuidtob62
  ```

- Output to a file:
  ```bash
  $ uuidtob62 input_file.txt > output_file.txt
  ```
