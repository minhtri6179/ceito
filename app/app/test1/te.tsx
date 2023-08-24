"use client";
import * as React from 'react';
import Box from '@mui/material/Box';
import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import ImageGallery from "../components/ImageGallery";
import ListeningQuestion from "../components/ListeningQuestion";
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
            <Tab label="Item One" value="1" />
            <Tab label="Item Two" value="2" />
            <Tab label="Item Three" value="3" />
            <Tab label="Item Three" value="4" />
            <Tab label="Item Three" value="5" />
            <Tab label="Item Three" value="6" />
            <Tab label="Item Three" value="7" />
          </TabList>
        </Box>
        <TabPanel value="1"><ImageGallery /></TabPanel>
        <TabPanel value="2"><ListeningQuestion index={6} numofQuestion={25} /></TabPanel>
        <TabPanel value="3"><ListeningQuestion index={31} numofQuestion={39} /></TabPanel>
        <TabPanel value="4"><ListeningQuestion index={70} numofQuestion={30} /></TabPanel>
        <TabPanel value="5"><ListeningQuestion index={5} numofQuestion={10} /></TabPanel>
        <TabPanel value="6"><ListeningQuestion index={5} numofQuestion={10} /></TabPanel>
        <TabPanel value="7"><ListeningQuestion index={5} numofQuestion={10} /></TabPanel>
      </TabContext>
    </Box>
  );
}