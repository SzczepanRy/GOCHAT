const socket = new WebSocket("ws://localhost:3000/ws");

const ul = document.querySelector(".messages");
const input = document.querySelector(".inputBox");
const form = document.querySelector(".messageBar");

console.log(document.cookie);

function getToken() {
    let cookie = document.cookie
    let token = cookie.split(";")[0].split("=")[1]
    return token
}



window.addEventListener("load", () => {
    (async () => {


        let res = await fetch("/api/validate", {
            method: "POST",
            headers: { "Content-type": "application/json" }
            ,
            body: JSON.stringify({ accessToken: getToken() })
        })

        let data = await res.json()

        if (!data.ok) {

            window.location.href = "/"
        }
    })()
    socket.onopen = () => {
        const title = document.querySelector(".title")
        title.innerText = "CONNECTED";
    };
});





form.addEventListener("submit", (e) => {
    e.preventDefault();
    if (input.value) {
        socket.send(input.value);
        input.value = "";
    }
    input.focus();
});
socket.onmessage = (e) => {
    console.log(e.data);
    const li = document.createElement("li");
    li.innerText = e.data;

    ul.append(li);
};


const logoutBtn = document.querySelector(".logout")
logoutBtn.addEventListener("click", () => {
    console.log("clear")

    cookieStore.delete("token");
    window.location.href = "/"

})


const createChat = document.querySelector(".createChat")
const createChatButton = document.querySelector(".createChatButton")

createChatButton.addEventListener("click", () => {
    createChat.showModal();
})
const cancelChatButton = document.querySelector(".cancelChat")

cancelChatButton.addEventListener("click", () => {
    createChat.close();
})

const availibleChats = document.querySelector(".availibleChats")
const joinChatButton = document.querySelector(".joinChatButton")
joinChatButton.addEventListener("click", () => {
    availibleChats.showModal()
})

const closeAvailible = document.querySelector(".closeAvailible")
closeAvailible.addEventListener("click", () => {
    availibleChats.close()
})
