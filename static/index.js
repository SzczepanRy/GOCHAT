const socket = new WebSocket("ws://localhost:3000/ws");

const ul = document.querySelector(".messages");
const input = document.querySelector(".inputBox");
const form = document.querySelector(".messageBar");

console.log("js works");

window.addEventListener("load", () => {
    socket.onopen = () => {
        const title= document.querySelector(".title")
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


const createChat = document.querySelector(".createChat")
const createChatButton = document.querySelector(".createChatButton")

createChatButton.addEventListener("click",()=>{
    createChat.showModal();
})
const cancelChatButton = document.querySelector(".cancelChat")

cancelChatButton.addEventListener("click",()=>{
    createChat.close();
})

const availibleChats = document.querySelector(".availibleChats")
const joinChatButton = document.querySelector(".joinChatButton")
joinChatButton.addEventListener("click",()=>{
    availibleChats.showModal()
})

const  closeAvailible = document.querySelector(".closeAvailible")
closeAvailible.addEventListener("click",()=>{
    availibleChats.close()
})
