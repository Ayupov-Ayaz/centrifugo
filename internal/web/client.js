

const host = "localhost:8000";
const connectionPath = "/connection/websocket";
const getTokenPath = "/centrifugo/connection_token";
const ws = "ws://";
const http = "http://";

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlRvbW15IiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.bA_T0Z34c0I84bcJ0Hgiaaq-f3z2TfKkdM9uR6tvLAU";

async function getToken() {
    console.log("getToken");

    if (!loggedIn) {
        return ""; // Empty token or pre-generated token for anonymous users.
    }
    // Fetch your application backend.
    const res = await fetch('http://'+host+'/centrifugo/connection_token');
    if (!res.ok) {
        if (res.status === 403) {
            // Return special error to not proceed with token refreshes,
            // client will be disconnected.
            throw new Centrifuge.UnauthorizedError();
        }
        // Any other error thrown will result into token refresh re-attempts.
        throw new Error(`Unexpected status code ${res.status}`);
    }
    const data = await res.json();
    console.log(data);
    return data.token;
}

const client = new Centrifuge(
    'ws://'+host+'/connection/websocket',
    {
        token: token, // Optional, getToken is enough.
        getToken:  getToken
    }
);

function display(message, data) {
    console.log(message, data);
}

client.on("connect", function (ctx) {
    display("connected", ctx);
});

client.on("disconnect", function (ctx) {
    display("disconnected", ctx);
});

client.on("closed", function (ctx) {
    display("closed", ctx);
})

client.on("error", function (err) {
    display("error", err);
});

client.connect();
console.log(client)