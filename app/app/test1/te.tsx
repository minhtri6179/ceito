"use client";
import * as React from 'react';
import Box from '@mui/material/Box';
import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import Button from '@mui/material/Button';
import TabPanel from '@mui/lab/TabPanel';
import ImageGallery from "../components/ImageGallery";
import ListeningQuestion from "../components/ListeningQuestion2";
import ListeningQuestionTest345 from "../components/ListeningQuestion3_4"
import PublishIcon from '@mui/icons-material/Publish';
import SubmitButton from "../components/Submit";
export default function LabTabs() {
  const [value, setValue] = React.useState('1');

  const handleChange = (event: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
  };

  return (
    <Box sx={{ width: '100%', typography: 'body1' }}>
      
      <TabContext value={value}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <TabList onChange={handleChange} aria-label="lab API tabs example">
            <Tab label="Part 1" value="1" />
            <Tab label="Part 2" value="2" />
            <Tab label="Part 3" value="3" />
            <Tab label="Part 4" value="4" />
            <Tab label="Part 5" value="5" />
            <Tab label="Part 6" value="6" />
            <Tab label="Part 7" value="7" />
          </TabList>
        </Box>
        <TabPanel value="1"><ImageGallery /></TabPanel>
        <TabPanel value="2"><ListeningQuestion index={6} numofQuestion={25} /></TabPanel>
        <TabPanel value="3"><ListeningQuestionTest345 name_part={"ETS-23-Test1-Part3"} /></TabPanel>
        <TabPanel value="4"><ListeningQuestionTest345 name_part={"ETS-23-Test1-Part4"} /></TabPanel>
        <TabPanel value="5"><ListeningQuestionTest345 name_part={"ETS-23-Test1-Part5"} /></TabPanel>
        <TabPanel value="6"><ListeningQuestionTest345 name_part={"ETS-23-Test1-Part6"} /></TabPanel>
        <TabPanel value="7"><ListeningQuestionTest345 name_part={"ETS-23-Test1-Part7"} /></TabPanel>
        <Button
          component="label"
          variant="contained"
          startIcon={<PublishIcon />}
          style={{
            position: 'fixed',
            right: '20%', // Adjust this value to control the button's distance from the right edge
            bottom: '70%', // Adjust this value to control the button's distance from the bottom
          }}
        >
          <SubmitButton />
        </Button>
      </TabContext>
      
    </Box>
  );
}