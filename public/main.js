
async function checkPings() {
    const counter = document.getElementById("pingCount");
    setInterval(async () => {
        const res = await fetch("/ping");
        if (!res.ok) {
            console.log(res);
            return
        }

        const data = await res.json();
        counter.textContent = data.count ?? '0';
    }, 1000)
}

function start() {
    checkPings();
}

start();
