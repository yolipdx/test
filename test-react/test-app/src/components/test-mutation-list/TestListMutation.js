import { useState } from "react";

let nextId = 3;
const initialList = [
    { id: 0, title: 'Big Bellies', seen: false },
    { id: 1, title: 'Lunar Landscape', seen: false },
    { id: 2, title: 'Terracotta Army', seen: true },
];


function ItemList({ artworks, onToggle }) {
    return (
      <ul>
        {artworks.map(artwork => (
          <li key={artwork.id}>
            <label>
              <input
                type="checkbox"
                checked={artwork.seen}
                onChange={e => {
                  onToggle(
                    artwork.id,
                    e.target.checked
                  );
                }}
              />
              {artwork.title}
            </label>
          </li>
        ))}
      </ul>
    );
  }

export default function BucketList() {
    const [myList, setMyList] = useState(initialList);
    const [yourList, setYourList] = useState(initialList);

    function handleToggleMyList(artworkId, nextSeen) {
        const newList = myList.map(artWork => {
            if (artWork.id === artworkId) {
                return {
                    ...artWork,
                    seen: !artWork.seen,
                }
            }

            return artWork;
        });

        setMyList(newList);
    }

    function handleToggleYourList(artworkId, nextSeen) {
        const newList = yourList.map(artWork => {
            if (artWork.id === artworkId) {
                return {
                    ...artWork,
                    seen: !artWork.seen,
                }
            }

            return artWork;
        });

        setYourList(newList);
    }

    return (
        <>
          <h1>Art Bucket List</h1>
          <h2>My list of art to see:</h2>
          <input />
          <br />
          <input size={21} />
          <br />
          <input size={22} />
          <br />
          <input size={23} />
          <br />
          <input size={24} />
          <br />
          <input size={25} />
          <ItemList
            artworks={myList}
            onToggle={handleToggleMyList} />
          <h2>Your list of art to see:</h2>
          <ItemList
            artworks={yourList}
            onToggle={handleToggleYourList} />
        </>
    );

}