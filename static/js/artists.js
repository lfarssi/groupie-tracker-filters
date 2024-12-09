const artistContainer = document.getElementById("artistContainer");
const suggestion = document.getElementById("suggestions");
const locationSuggestion = document.getElementById("suggestions_location");
// Get artist data from the embedded JSON script tag asynchronously
const artistDataScript = document.getElementById("artistData");
const artists = JSON.parse(artistDataScript.textContent);
const searchArtist = [...artists]
const uniqueLocations = new Set();
const uniqueDates = new Set();  
console.log(artists);
searchArtist.forEach((artist) => {
artist.Locationsr.forEach((city) => {
  if (!uniqueLocations.has(city)) {
    uniqueLocations.add(city);
    const locations = document.createElement("option");
    locations.value = city;
    locationSuggestion.appendChild(locations);
  }
});
} );
searchArtist.forEach((artist) => {
  const artistName = document.createElement("option");
  artistName.value = artist.name;
  artistName.innerHTML = "- Artist/Band";
  suggestion.appendChild(artistName);

  const firstAlbum = document.createElement("option");
  firstAlbum.value = artist.firstAlbum;
  firstAlbum.innerHTML = "- First Album";
  suggestion.appendChild(firstAlbum);

  const creationDate = document.createElement("option");
  creationDate.value = artist.creationDate;
  creationDate.innerHTML = "- Creation Date";
  suggestion.appendChild(creationDate);

  artist.members.forEach((member) => {
    const members = document.createElement("option");
    members.value = member;
    members.innerHTML = "- Member";
    suggestion.appendChild(members);
  });

  artist.Locationsr.forEach((city) => {
    if (!uniqueLocations.has(city)) {
      uniqueLocations.add(city);
      const locations = document.createElement("option");
      locations.value = city;
      locations.innerHTML = "- Location";
      suggestion.appendChild(locations);
    }
  });

 


  artist.ConcertDatesr.forEach((date) => {
    if (!uniqueDates.has(date)) {
      uniqueDates.add(date);
      const dates = document.createElement("option");
      dates.value = date;
      dates.innerHTML = "- Date";
      suggestion.appendChild(dates);
    }
  });
});

// Render artists in the container
artistContainer.innerHTML = artists
  .map(
    (artist) => `
                <div class="card">
                    <a href="/artist/${artist.id}">
                    <div class="div-image">
                        <img src="${artist.image}" alt="${artist.name}">
                    </div>
                        <h2>${artist.name}  </h2>
                        <p>${artist.Type}</p>
                    </a>
                </div>
            `
  )
  .join("");
