const loginNav = document.querySelector(".loginNav")
const redgisterNav = document.querySelector(".redgisterNav")

const loginForm = document.querySelector(".loginForm")
const redgisterForm = document.querySelector(".redgisterForm")


redgisterNav.addEventListener("click", () => {
    loginForm.style.display = "none"

    redgisterForm.style.display = "flex"
})
loginNav.addEventListener("click", () => {
    loginForm.style.display = "flex"

    redgisterForm.style.display = "none"
})



const loginInputL = document.querySelector(".loginL")
const passwordInputL = document.querySelector(".passwordL")

const loginButton = document.querySelector(".loginButton")

loginButton.addEventListener("click", () => {
    let login = loginInputL.value
    let password = passwordInputL.value
    if (login != "" && password != "") {
        (async () => {
            let res = await fetch("/api/login", {
                method: "POST",
                headers: { "content-type": "application/json" },
                body: JSON.stringify({ login, password })
            })
            let data = await res.json()
            if (data.success) {
                //console.log("adding token")
                //localStorage.setItem("accessToken", data.accessToken)

                const d = new Date()
                d.setTime(d.getTime() + (1000* 60*60))

                console.log(login)

                document.cookie=`token=${data.accessToken};expires=${d.toUTCString()};path=/`

                document.cookie=`login=${login};expires=${d.toUTCString()};path=/`
                console.log(document.cookie)

                window.location.href = "/static"

            }

            console.log(data)
        })()
    }
    else {
        console.log("did not fill out boulth fealds")
    }

})



const loginInputR = document.querySelector(".loginR")
const passwordInputR = document.querySelector(".passwordR")

const redgisterButton = document.querySelector(".redgisterButton")

redgisterButton.addEventListener("click", () => {
    let login = loginInputR.value
    let password = passwordInputR.value
    if (login != "" && password != "") {
        (async () => {
            let res = await fetch("/api/redgister", {
                method: "POST",
                headers: { "content-type": "application/json" },
                body: JSON.stringify({ login, password })
            })
            let data = await res.json()
            console.log(data)

        })()
    }
    else {
        console.log("did not fill out boulth fealds")
    }
})

