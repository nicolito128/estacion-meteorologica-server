const counter = document.getElementById("ping");

window.onload = start;

async function start() {
    setIntervalPings();
}

async function setIntervalPings() {
    updatePing();
    window.pingIntervalId = setInterval(updatePing, 1000);
}

async function updatePing() {
    try {
        const res = await fetch("/ping");
        if (!res.ok) {
            console.log(res);
            return
        }

        const data = await res.json();
        if (!data) {
            console.log(data);
            return
        }

        counter.textContent = parseMillisecondsIntoReadableTime(data.ping);
    } catch (e) {
        console.error(e);
        clearInterval(window.pingIntervalId);
    }
}

function parseMillisecondsIntoReadableTime(milliseconds) {
    const hours = milliseconds / (1000 * 60 * 60);
    const absoluteHours = Math.floor(hours);
    const h = absoluteHours > 9 ? absoluteHours : '0' + absoluteHours;

    const minutes = (hours - absoluteHours) * 60;
    const absoluteMinutes = Math.floor(minutes);
    const m = absoluteMinutes > 9 ? absoluteMinutes : '0' + absoluteMinutes;

    const seconds = (minutes - absoluteMinutes) * 60;
    const absoluteSeconds = Math.floor(seconds);
    const s = absoluteSeconds > 9 ? absoluteSeconds : '0' + absoluteSeconds;

    return h + ':' + m + ':' + s;
}

