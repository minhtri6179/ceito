import * as React from "react";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";

export default function ContainedButtons() {
  return (
    <Stack direction="row" spacing={2}>
      <Button variant="contained">Part1</Button>
      <Button variant="outlined">Part2</Button>
      <Button variant="outlined">Part3</Button>
      <Button variant="outlined">Part4</Button>
      <Button variant="outlined">Part5</Button>
      <Button variant="outlined">Part6</Button>
      <Button variant="outlined">Part7</Button>
    </Stack>
  );
}
