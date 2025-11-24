// index.js
import * as pdfjsLib from '/assets/js/pdf.mjs'; // Adjust path as needed

pdfjsLib.GlobalWorkerOptions.workerSrc = '/assets/js/pdf.worker.mjs'; // Adjust path as needed

export function reload(scaleXY) {
  const canvas = document.getElementById('pdf-canvas');
  const context = canvas.getContext('2d');
  const pdfUrl = 'pdf.pdf'; // Replace with your PDF file

  pdfjsLib.getDocument(pdfUrl).promise.then(function (pdfDoc) {
    // Load the first page
    pdfDoc.getPage(1).then(function (page) {
      const scale = scaleXY; // Adjust scale as needed
      const viewport = page.getViewport({ scale: scale });

      canvas.height = viewport.height;
      canvas.width = viewport.width;
      const renderContext = {
        canvasContext: context,
        viewport: viewport
      };

      page.render(renderContext).promise.then(function () {
        console.log('Scale ', scaleXY);
        console.log('Scale viewport', viewport);
        console.log('Page rendered successfully!');
      });
    });
  }).catch(function (error) {
    console.error('Error loading PDF:', error);
  });
}
