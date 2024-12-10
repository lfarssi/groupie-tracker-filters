const artistContainer = document.getElementById("artistContainer");
const suggestion = document.getElementById("suggestions");
const locationSuggestion = document.getElementById("suggestions_location");
// Get artist data from the embedded JSON script tag asynchronously 
const artistDataScript = document.getElementById("artistData");
const suggestionsDataScript = document.getElementById("suggetionsData");
const artists = JSON.parse(artistDataScript.textContent);
const suggests = JSON.parse(suggestionsDataScript.textContent);
const uniqueLocations = new Set();
const uniqueDates = new Set(); 
const uniqueLocationsFilter = new Set()

console.log(artists);

suggests.forEach((suggest) => {
  const artistName = document.createElement("option");
  artistName.value = suggest.name;
  artistName.innerHTML = "- Artist/Band";
  suggestion.appendChild(artistName);

  const firstAlbum = document.createElement("option");
  firstAlbum.value = suggest.firstAlbum;
  firstAlbum.innerHTML = "- First Album";
  suggestion.appendChild(firstAlbum);

  const creationDate = document.createElement("option");
  creationDate.value = suggest.creationDate;
  creationDate.innerHTML = "- Creation Date";
  suggestion.appendChild(creationDate);

  suggest.members.forEach((member) => {
    const members = document.createElement("option");
    members.value = member;
    members.innerHTML = "- Member";
    suggestion.appendChild(members);
  });

  suggest.Locationsr.forEach((city) => {
    if (!uniqueLocations.has(city)) {
      uniqueLocations.add(city);
      const locations = document.createElement("option");
      locations.value = city;
      locations.innerHTML = "- Location";
      suggestion.appendChild(locations);
    }
  });

 


  suggest.ConcertDatesr.forEach((date) => {
    if (!uniqueDates.has(date)) {
      uniqueDates.add(date);
      const dates = document.createElement("option");
      dates.value = date;
      dates.innerHTML = "- Date";
      suggestion.appendChild(dates);
    }
  });
});

suggests.forEach((suggest) => {
  suggest.Locationsr.forEach((city) => {
    if (!uniqueLocationsFilter.has(city)) {
      uniqueLocationsFilter.add(city);
      const locations = document.createElement("option");
      locations.value = city;
      locations.textContent = "- Location";
      locationSuggestion.appendChild(locations);
    }
  });
  } );
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
  if (artists.length == 0) {
    artistContainer.innerHTML = `
    <div class="no-results">
        <h2>No results found</h2>
        <p>Try to search for another artist</p>
    </div>
    `;
  }
