import submit from "../../util/Submit";
import  React, { useState } from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';
const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};
const SubmitButton = () => {
  const [result, setResult] = useState(""); // Initialize result with an empty string
  const [open, setOpen] = React.useState(false);
  const handleClose = () => setOpen(false);


  const handleClick = () => {
    setOpen(true)
    submit()
      .then((response) => { // Capture the response from submit()
        setResult(response); // Set the result with the response data
      })
      .catch((error) => {
        console.error("Error submitting data:", error);
      });
  };

  return (
    <div>
      <button onClick={handleClick}><strong>Nộp bài thôi </strong></button>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            <strong>Kết quả nè: </strong>{result}
          </Typography>
        </Box>
      </Modal>
    </div>
  );
};

export default SubmitButton;
