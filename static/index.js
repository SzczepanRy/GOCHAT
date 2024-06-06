const socket = new WebSocket("ws://localhost:3000/ws");

const ul = document.querySelector(".messages");
const input = document.querySelector(".inputBox");
const form = document.querySelector(".messageBar");

console.log(document.cookie);

function getCookie(index) {
    let cookie = document.cookie

    let token = cookie.split(";")[index].split("=")[1]

    return token
}



window.addEventListener("load", () => {
    (async () => {


        let res = await fetch("/api/validate", {
            method: "POST",
            headers: { "Content-type": "application/json" }
            ,
            body: JSON.stringify({ accessToken: getCookie(1) })
        })

        let data = await res.json()

        if (!data.ok) {

            window.location.href = "/"
        }
        document.querySelector(".loginInfo").innerText ="logged in as " + getCookie(0)

        //display chats


        res = await fetch("/api/getChatRooms", {
            method: "GET",
            headers: { "Content-type": "application/json" }
            ,
        })

        data = await res.json()

        const chatList = document.querySelector(".chatList")
        data.chatNames.map((name)=>{
            let li = document.createElement("li")
            li.innerText = name
            chatList.append(li)

        })



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
