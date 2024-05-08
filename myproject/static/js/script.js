const menuData = [
    {id: 1, name: "Item 1", price: 10},
    {id: 2, name: "Item 2", price: 15},
    {id: 3, name: "Item 3", price: 20}
];

const cart = [];

function renderMenu() {
    const menuContainer = document.getElementById("menu");

    menuData.forEach(item => {
        const itemElement = document.createElement("div");
        itemElement.classList.add("item");
        itemElement.innerHTML = `
            <h3>${item.name}</h3>
            <p>Price: $${item.price}</p>
            <button onclick="addToCart(${item.id})">Add to Cart</button>
        `;
        menuContainer.appendChild(itemElement);
    });
}

function addToCart(itemId) {
    const item = menuData.find(item => item.id === itemId);
    if (item) {
        cart.push(item);
        renderCart();
    }
}


    function redirectTo(url) {
    window.location.href = url + ".html";
}


function renderCart() {
    const cartItemsContainer = document.getElementById("cart-items");
    const totalElement = document.getElementById("total");
    let totalPrice = 0;

    cartItemsContainer.innerHTML = "";

    cart.forEach(item => {
        const cartItemElement = document.createElement("li");
        cartItemElement.innerHTML = `${item.name} - $${item.price}`;
        cartItemElement.innerHTML += `<button onclick="removeFromCart(${item.id})">Remove</button>`;
        cartItemsContainer.appendChild(cartItemElement);
        totalPrice += item.price;
    });

    totalElement.textContent = totalPrice.toFixed(2);
}

function removeFromCart(itemId) {
    const itemIndex = cart.findIndex(item => item.id === itemId);
    if (itemIndex !== -1) {
        cart.splice(itemIndex, 1);
        renderCart();
    }
}

renderMenu();
