import * as React from "react";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import Link from "next/link";
export default function Navbar() {
  return (
    <Stack direction="row" spacing={2}>
      <Link href="/test1/part1">
        <Button variant="contained">Part1</Button>
      </Link>
      <Link href="/test1/part2">
        <Button variant="outlined">Part2</Button>
      </Link>

      <Link href="/test1/part3">
        <Button variant="outlined">Part3</Button>
      </Link>

      <Link href="/test1/part4">
        <Button variant="outlined">Part4</Button>
      </Link>

      <Link href="/test1/part5">
        <Button variant="outlined">Part5</Button>
      </Link>

      <Link href="/test1/part6">
        <Button variant="outlined">Part6</Button>
      </Link>

      <Link href="/test1/part7">
        <Button variant="outlined">Part7</Button>
      </Link>
    </Stack>
  );
}
