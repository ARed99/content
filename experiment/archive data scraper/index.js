const puppeteer = require('puppeteer');
const fs = require('fs');

(async () => {
  // Launch a new browser instance
  const browser = await puppeteer.launch({headless:false});

  // Open a new page
  const page = await browser.newPage();

  // Navigate to the desired webpage
  await page.goto('https://archive.org/details/texts?&sort=-week&page=12'); // Replace with your target URL
 await page.screenshot('img.png' )
  // Wait for the elements to load on the page

  
  // Extract the href attributes of the <a> tags
    await page.$$eval('.cell-container', (elements) =>
    console.log(elements)
  );
  
  // Log the href values


  // Close the browser
  await browser.close();
})();
