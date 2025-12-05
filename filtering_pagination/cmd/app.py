import streamlit as st
import requests
import math
import html
import fashion  # your python file

API_URL = "http://localhost:8080/FashionStore/v1/filter"

st.set_page_config(page_title="E-Com Store", layout="wide")

# Load from fashion.py
CATEGORIES = fashion.Category()
BRANDS = fashion.Brand()
GENDERS = fashion.Gender()
COLORS = fashion.Color()
SIZES = fashion.Size()

# ---------------- CSS -----------------
st.markdown("""
<style>
.ecom-card {
  border: 1px solid #ddd;
  border-radius: 12px;
  padding: 10px;
  background: white;
  margin-bottom: 15px;
}
.ecom-title { font-size: 16px; font-weight: 600; margin: 5px 0; }
.ecom-brand { color: grey; font-size: 13px; margin: 2px 0; }
.ecom-price { font-weight: bold; color: #111; font-size: 18px; margin: 2px 0; }
.ecom-rating { color: #f59e0b; font-size: 14px; }
</style>
""", unsafe_allow_html=True)

# ---------------- TITLE ----------------
st.title("🛍️ E-Commerce Store")

# ---------------- FILTERS ----------------
st.sidebar.header("Filters")

selected_categories = st.sidebar.multiselect(
    "Categories",
    options=list(CATEGORIES.keys()),
    format_func=lambda k: CATEGORIES[k]
)

selected_brands = st.sidebar.multiselect("Brands", BRANDS)
selected_colors = st.sidebar.multiselect("Colors", COLORS)
selected_sizes = st.sidebar.multiselect("Sizes", SIZES)
selected_genders = st.sidebar.multiselect("Gender", GENDERS)

min_price, max_price = st.sidebar.slider("Price Range", 0, 20000, (0, 20000))
min_rating = st.sidebar.slider("Minimum Rating", 0.0, 5.0, 0.0, 0.1)

# Pagination state
if "page" not in st.session_state:
    st.session_state.page = 1

page = st.session_state.page
limit = 12  # 3×4 layout

# ---------------- REQUEST PARAMS (ARRAY SUPPORT) ----------------
params = {
    "min_price": min_price,
    "max_price": max_price,
    "rating": min_rating,
    "page": page,
    "limit": limit,
}

# Arrays — requests automatically encodes as `brand=Nike&brand=Puma`
if selected_categories:
    params["category_id"] = selected_categories

if selected_brands:
    params["brand"] = selected_brands

if selected_colors:
    params["color"] = selected_colors

if selected_sizes:
    params["size"] = selected_sizes

if selected_genders:
    params["gender"] = selected_genders

# Debug print to terminal
print("\n========= API REQUEST PARAMS =========")
print(params)
print("=====================================\n")

# ---------------- API CALL ----------------
try:
    resp = requests.get(API_URL, params=params)
    resp.raise_for_status()
    data = resp.json()
except Exception as e:
    st.error(f"API error: {e}")
    st.stop()

# Normalize list response
products = data if isinstance(data, list) else []
total = len(products)

st.subheader(f"Showing {total} products")

# ---------------- GRID VIEW ----------------
cols_per_row = 3
rows = math.ceil(total / cols_per_row)

idx = 0
for r in range(rows):
    cols = st.columns(cols_per_row)
    for col in cols:
        if idx >= total:
            break

        p = products[idx]

        # helper for fields
        get = lambda *k: next((p.get(x) for x in k if x in p), None)

        title = get("Name", "name") or "Unnamed"
        brand = get("Brand", "brand") or ""
        price = get("BasePrice", "base_price") or "-"
        rating = get("Rating", "rating") or "-"

        img = f"https://picsum.photos/seed/{idx}/400/400"

        col.markdown(
            f"""
            <div class="ecom-card">
                <img src="{img}" width="100%" style="border-radius:8px;">
                <div class="ecom-title">{title}</div>
                <div class="ecom-brand">{brand}</div>
                <div class="ecom-price">₹{price}</div>
                <div class="ecom-rating">⭐ {rating}</div>
            </div>
            """,
            unsafe_allow_html=True
        )
        idx += 1

# ---------------- PAGINATION ----------------
prev_col, page_col, next_col = st.columns([1, 2, 1])

with prev_col:
    if page > 1 and st.button("⬅ Previous"):
        st.session_state.page -= 1
        st.rerun()

with page_col:
    st.write(f"Page {page}")

with next_col:
    if total == limit and st.button("Next ➡"):
        st.session_state.page += 1
        st.rerun()
