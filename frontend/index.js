console.log("js loading");

document.addEventListener('DOMContentLoaded', () => {
    console.log("loaded");
    const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjAwODQwMjQsImlhdCI6MTcxOTk5NzYyNCwidXNlcm5hbWUiOiJIYW1tYWRBaG1hZCJ9.wJZ9QyWDtWNv8Q3mybjva8udYBwTae9-WXx57KhlCHc'
    const socket = new WebSocket(`ws://localhost:1323/ws?token=${token}`);

    socket.onopen = function (event) {
        console.log('WebSocket connected');
    };

    socket.onmessage = function (event) {
        const data = JSON.parse(event.data);
        console.log('Received data from server:', data);

        updateTable(data.data);
    };

    socket.onclose = function (event) {
        console.log('WebSocket closed');
    };


    let button = document.getElementById('upBtn');
    console.log(button);

    button.addEventListener('click', (e) => {
        handleFileUpload(e);
    });
    let goroutinesInput = document.getElementById('GoRoutines');
    console.log(goroutinesInput);
    button.addEventListener('click', (e) => {
        console.log("working");
    });

    function updateTable(data) {

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

    function handleFileUpload(event) {
        const file = fileInput.files[0];
        if (!file) {
            alert('No file selected');
            return;
        }

        const formData = new FormData();
        formData.append('file', file);
        formData.append('goroutines', goroutinesInput.value)


        fetch('http://localhost:1323/textfileprocessor', {
            method: 'POST',
            headers: {

                'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjAwODQwMjQsImlhdCI6MTcxOTk5NzYyNCwidXNlcm5hbWUiOiJIYW1tYWRBaG1hZCJ9.wJZ9QyWDtWNv8Q3mybjva8udYBwTae9-WXx57KhlCHc'
            },
            body: formData
        })
            .then(response => response.json())
            .then(data => {
                console.log('File upload response:', data);

            })
            .catch(error => {
                console.error('Error uploading file:', error);
            });
    }
});
