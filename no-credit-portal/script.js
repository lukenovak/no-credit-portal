function submitSongForm() {
    let xmlHttpRequest = new XMLHttpRequest();

    // if the response is 200 we show the success element
    xmlHttpRequest.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            document.getElementById("songSuccessMessage").classList.remove("d-none")
        }
    };

    let songName = document.getElementById("songNameInput").value;
    let artistName = document.getElementById("artistNameInput").value;

    if (songName === "" || artistName === "") {
        return;
    }

    // build the request and send it
    xmlHttpRequest.open('POST', "/submit/song", true);
    xmlHttpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xmlHttpRequest.send("song=" + songName + "&artist=" + artistName);
}

function submitGBPEntry() {
    let xmlHttpRequest = new XMLHttpRequest();

    let itemName = document.getElementById("gbpInput").value;

    console.log("this part of the api isn't set up yet but your item is: " + itemName);

}