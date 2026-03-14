# Week 1 Checklist

## Weekly Targets

* [ ] Solve **21 LeetCode problems**
* [ ] Watch **7 Arpit Bhayani system design videos**
* [ ] Read **2 engineering blogs**
* [ ] Build **Go URL Shortener Backend**
* [ ] Write **Sunday reflection**

---

# System Design (Office Hours)

### Arpit Bhayani Videos

* [ ] Monday
* [ ] Tuesday
* [ ] Wednesday
* [ ] Thursday
* [ ] Friday
* [ ] Saturday
* [ ] Sunday

---

### Engineering Blogs

* [ ] **Discord — How Discord indexes billions of messages**
  [https://discord.com/blog/how-discord-indexes-billions-of-messages](https://discord.com/blog/how-discord-indexes-billions-of-messages)

* [ ] **Instagram — Video upload latency improvements**
  [https://instagram-engineering.com/video-upload-latency-improvements-at-instagram-bcf4b4c5520a](https://instagram-engineering.com/video-upload-latency-improvements-at-instagram-bcf4b4c5520a)

---

# LeetCode (9–10 PM)

## Monday

* [ ] Two Sum
* [ ] Largest Element in Array
* [ ] Second Largest Element in Array

---

## Tuesday

* [ ] Check if Array is Sorted
* [ ] Remove Duplicates from Sorted Array
* [ ] Left Rotate Array by One

---

## Wednesday

* [ ] Move Zeroes
* [ ] Union of Two Sorted Arrays
* [ ] Missing Number

---

## Thursday

* [ ] Maximum Consecutive Ones
* [ ] Single Number
* [ ] Longest Subarray With Sum K

---

## Friday

* [ ] Kadane's Algorithm
* [ ] Rearrange Array by Sign
* [ ] Best Time to Buy and Sell Stock

---

## Saturday

* [ ] Next Permutation
* [ ] Majority Element
* [ ] Set Matrix Zeroes

---

## Sunday

* [ ] Pascal's Triangle
* [ ] Rotate Matrix
* [ ] Merge Intervals

---

# Go Backend Project (10–11 PM)

## Project: URL Shortener API

---

## Monday — Project Setup

* [ ] Initialize Go module
* [ ] Create folder structure

```
cmd/
internal/
handlers/
services/
repository/
models/
configs/
```

* [ ] Create basic HTTP server
* [ ] Add `/health` endpoint

---

## Tuesday — Shorten API

* [ ] Create `POST /shorten`
* [ ] Generate short ID
* [ ] Store mapping in memory map

Example request:

```
POST /shorten
{
"url":"https://example.com"
}
```

---

## Wednesday — Redirect API

* [ ] Create `GET /{shortID}`
* [ ] Lookup original URL
* [ ] Implement HTTP redirect

---

## Thursday — Database

* [ ] Integrate SQLite
* [ ] Create URL table
* [ ] Store shortened links in DB

---

## Friday — Code Structure

* [ ] Separate handlers / services / repository
* [ ] Move business logic to service layer
* [ ] Clean up project structure

---

## Saturday — Analytics

* [ ] Track click counts
* [ ] Increment counter on redirect
* [ ] Create endpoint

```
GET /stats/{shortID}
```

---

## Sunday — Final Polish

* [ ] Refactor code
* [ ] Test all endpoints
* [ ] Add README for project
* [ ] Push to GitHub

---

# Sunday Reflection

Fill this at end of week.

* [ ] Problems solved: ___ / 21
* [ ] Hardest problem:
* [ ] Most interesting Go concept learned:
* [ ] Best system design insight:
* [ ] Improvements for next week:

---

# Non-Negotiable Rule

```
Never break the chain.
```

Even on a bad day:

* [ ] Solve **at least 1 problem**
