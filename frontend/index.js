console.log("js loading");
document.addEventListener('DOMContentLoaded', () => {
    console.log("loaded");
    const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjAwNjk2ODcsImlhdCI6MTcxOTk4MzI4NywidXNlcm5hbWUiOiJNdWhhbW1hZEFobWFkIn0.kcBEiux_R2Mf2yNGGteaxhBsRjz3EwVBE1JnOankEbQ'
    const socket = new WebSocket(`ws://localhost:1323/ws?token=${token}`);

    socket.onopen = function (event) {
        console.log('WebSocket connected');
    };

    socket.onmessage = function (event) {
        const data = JSON.parse(event.data);
        console.log('Received data from server:', data);

        // Update UI with received data (e.g., display in table)
        updateTable(data.data);
    };

    socket.onclose = function (event) {
        console.log('WebSocket closed');
    };
    //upload button would not trigger
    // Select all elements with the class 'upload-button'

    let button = document.getElementById('upBtn');
    console.log(button);
    // Iterate over each button element
    // for (let button of buttons) {
    //     // Add a click event listener to each button
    //     button.addEventListener('click', (e) => {
    //         handleFileUpload(e); // Call handleFileUpload function when button is clicked
    //     });
    // }
    button.addEventListener('click', (e) => {
        handleFileUpload(e); // Call handleFileUpload function when button is clicked
    });
    let goroutinesInput = document.getElementById('GoRoutines');
    console.log(goroutinesInput);
    button.addEventListener('click', (e) => {
        console.log("working"); // Call handleFileUpload function when button is clicked
    });

    function updateTable(data) {
        // Update the HTML table with the received data
        const tableBody = document.getElementById('tableBody');
        tableBody.innerHTML = `
            <tr>
                <td>${data.id}</td>
                <td>${data.totallines}</td>
                <td>${data.totalwords}</td>
                <td>${data.totalspaces}</td>
                <td>${data.totalvowels}</td>
                <td>${data.totalpunctuation}</td>
                <td>${data.timestamp}</td>
            </tr>
        `;
    }

    const fileInput = document.getElementById('fileInput');
    // fileInput.addEventListener('change', handleFileUpload);
    // this is responsible for uploading file
    function handleFileUpload(event) {
        const file = fileInput.files[0];
        if (!file) {
            alert('No file selected');
            return;
        }

        const formData = new FormData();
        formData.append('file', file);
        formData.append('goroutines', goroutinesInput.value)

        // Example: Send file via fetch to backend API
        fetch('http://localhost:1323/textfileprocessor', {
            method: 'POST',
            headers: {
                // Add headers as needed (e.g., Authorization for JWT)
                'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjAwNjk2ODcsImlhdCI6MTcxOTk4MzI4NywidXNlcm5hbWUiOiJNdWhhbW1hZEFobWFkIn0.kcBEiux_R2Mf2yNGGteaxhBsRjz3EwVBE1JnOankEbQ'
            },
            body: formData
        })
            .then(response => response.json())
            .then(data => {
                console.log('File upload response:', data);
                // Assuming backend sends data in the format specified
                // You can update UI with received data here if needed
            })
            .catch(error => {
                console.error('Error uploading file:', error);
            });
    }
});
