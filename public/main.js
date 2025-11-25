
window.onload = start;

async function start() {
    const humidityData = {
        name: "humidity",
        path: "/measurements/humidity",
        interval: 10000,
        displayNode: document.getElementById("rhBody"),
        avgNode: document.getElementById("rhAvg"),
        minNode: document.getElementById("rhMin"),
        maxNode: document.getElementById("rhMax"),
    };

    const temperatureData = {
        name: "temperature",
        path: "/measurements/temperature",
        interval: 10000,
        displayNode: document.getElementById("tempBody"),
        avgNode: document.getElementById("tempAvg"),
        minNode: document.getElementById("tempMin"),
        maxNode: document.getElementById("tempMax"),
    };

    handleInterval(updatePing, 1000);
    handleInterval(updateMeasurement(humidityData), humidityData.interval);
    handleInterval(updateMeasurement(temperatureData), temperatureData.interval);
}

function updateMeasurement(options) {
    if (!options?.name || !options?.path) {
        console.error("Invalid options for updateMeasurement");
        return;
    }
    if (!options?.interval || isNaN(options.interval) || options.interval < 1000) {
        console.error("Invalid interval for updateMeasurement, setting to 5000ms");
        options.interval = 5000;
    }
    if (!options?.displayNode || !options?.avgNode || !options?.minNode || !options?.maxNode) {
        console.error("Invalid display nodes for updateMeasurement");
        return;
    }

    return async function() {
        console.log("Updating measurement data for " + options.name);
        try {
            const result = await fetch(options.path).then(res => res.json()).catch(e => { throw e });
            if (!result?.data || result.data.length === 0) {
                throw new Error("No data received for " + options.name);
            }

            const samples = result.data;
            let html = "";
            let total = 0.0, minValue = samples[0]?.value ?? 9999, maxValue = samples[0]?.value ?? -9999;
            for (let i = samples.length - 1; i >= 0; i--) {
                const elem = samples[i];
                if (minValue > elem.value) minValue = elem.value;
                if (maxValue < elem.value) maxValue = elem.value;
                total += parseFloat(elem.value);
                html += `<tr><td>${elem.timestamp}</td> <td>${elem.host}</td> <td>${elem.value}</td></tr>`;
            }

            options.displayNode.innerHTML = html;
            options.avgNode.textContent = (total / samples.length).toFixed(2);
            options.minNode.textContent = minValue.toFixed(2);
            options.maxNode.textContent = maxValue.toFixed(2);
        } catch (e) {
            console.error("Error trying to update mesasurement data", e);
            clearInterval(window[`_updateMeasurement_${options.name}_interval`]);
        }
    }
}

async function handleInterval(callback, time) {
    callback();
    window[`_${callback.name}_interval`] = setInterval(callback, time);
}

async function updatePing() {
    const counter = document.getElementById("ping");
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

