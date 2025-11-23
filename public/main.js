const counter = document.getElementById("ping");

const temperatureDisplay = document.getElementById("tempBody");
const tempAvgElem = document.getElementById("tempAvg");
const tempMinElem = document.getElementById("tempMin");
const tempMaxElem = document.getElementById("tempMax");

const humidityDisplay = document.getElementById("rhBody");
const rhAvgElem = document.getElementById("rhAvg");
const rhMinElem = document.getElementById("rhMin");
const rhMaxElem = document.getElementById("rhMax");

window.onload = start;

async function start() {
    handleInterval(updateMeasurements, 5000);
    handleInterval(updatePing, 1000);
}

async function updateMeasurements() {
    console.log("Updating measurements...");
    try {
        const results = await Promise.all([
            fetch("/measurements/temperature").then(res => res.json()),
            fetch("/measurements/humidity").then(res => res.json()),
        ]).catch(e => { throw e });

        const temperatures = results[0]?.data;
        const humidities = results[1]?.data;

        if (temperatures && temperatures.length > 0) {
            let html = "";
            let total = 0.0;
            let minValue = temperatures[0]?.value ?? 999;
            let maxValue = temperatures[0]?.value ?? -999;
            for (let i = temperatures.length - 1; i >= 0; i--) {
                const elem = temperatures[i];
                if (minValue > elem.value) minValue = elem.value;
                if (maxValue < elem.value) maxValue = elem.value;
                total += elem.value;
                html += `<tr><td>${elem.timestamp}</td> <td>${elem.host}</td> <td>${elem.value}</td></tr>`;
            }
            temperatureDisplay.innerHTML = html;
            tempAvgElem.textContent = (total / temperatures.length).toFixed(2);
            tempMinElem.textContent = minValue.toFixed(2);
            tempMaxElem.textContent = maxValue.toFixed(2);
        }

        if (humidities && humidities.length > 0) {
            let html = "";
            let total = 0.0;
            let minValue = humidities[1]?.value ?? 999;
            let maxValue = humidities[1]?.value ?? -999;
            for (let i = humidities.length - 1; i >= 0; i--) {
                const elem = humidities[i];
                if (minValue > elem.value) minValue = elem.value;
                if (maxValue < elem.value) maxValue = elem.value;
                total += parseFloat(elem.value);
                html += `<tr><td>${elem.timestamp}</td> <td>${elem.host}</td> <td>${elem.value}</td></tr>`;
            }
            humidityDisplay.innerHTML = html;
            rhAvgElem.textContent = (total / humidities.length).toFixed(2);
            rhMinElem.textContent = minValue.toFixed(2);
            rhMaxElem.textContent = maxValue.toFixed(2);
        }

    } catch (e) {
        console.error(e);
        clearInterval(window[`_updateMeasurements_interval`]);
    }
}

async function handleInterval(callback, time) {
    callback();
    window[`_${callback.name}_interval`] = setInterval(callback, time);
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
        clearInterval(window[`_updatePing_interval`]);
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

