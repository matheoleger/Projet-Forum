function valueFromUrl() {
	url = window.location.href
    begin_url = url.match(/^.*(?=(\?))/gm).toString()

    perpage = url.match(/[?&]+perpage=?([^&]*)?/gi).toString()
    perpagev = perpage.match(/\d/gi).toString()
    perpagevalue = perpagev.replace(/,/g, '')

    page = url.match(/[?&]+page=?([^&]*)?/gi).toString()
    pagev = page.match(/\d/gi).toString()
    pagevalue = pagev.replace(/,/g, '')

    post = url.match(/[?&]+post=?([^&]*)?/gi)
    if (post != null ) {
        spost = post.toString()
        postv = spost.match(/\d/gi).toString()
        postvalue = postv.replace(/,/g, '')

        console.log("postvalue=" + postvalue + "perpagevalue=" + perpagevalue + "pagevalue=" + pagevalue )
        return [perpagevalue, pagevalue, begin_url, postvalue]
    }

    category = url.match(/[?&]+category=?([^&]*)?/gi)
    if (category != null ) {
        scategory = category.toString()
        categoryv = scategory.replace(/\?/g, '')
        categoryvalue = categoryv.replace(/category=/g, '')

        console.log("categoryvalue=" + categoryvalue + "perpagevalue=" + perpagevalue + "pagevalue=" + pagevalue )
        return [perpagevalue, pagevalue, begin_url, categoryvalue]
    }

    console.log("perpagevalue=" + perpagevalue + "pagevalue=" + pagevalue )
    return [perpagevalue, pagevalue, begin_url]
}

function nextPage() {
    values = valueFromUrl()

    page = parseInt(values[1]) + 1

    console.log("next page" + values)

    if (values[2] == "http://localhost:8080/posts/content") {
        window.location.replace("/posts/content?post=" + values[3] + "&perpage=" + values[0] + "&page=" + page)

    } else if (values[2] == "http://localhost:8080/posts") {
        window.location.replace("/posts?category=" + values[3] + "&perpage=" + values[0] + "&page=" + page)

    } else if (values[2] == "http://localhost:8080/categories") {
        window.location.replace("/categories?perpage=" + values[0] + "&page=" + page)

    } else {
        window.location.replace("/")
    }
}

function previousPage() {
    values = valueFromUrl()

    page = values[1]
    if (page > 0) {
        page--
    }

    console.log("next page" + values)

    if (values[2] == "http://localhost:8080/posts/content") {
        window.location.replace("/posts/content?post=" + values[3] + "&perpage=" + values[0] + "&page=" + page)

    } else if (values[2] == "http://localhost:8080/posts") {
        window.location.replace("/posts?category=" + values[3] + "&perpage=" + values[0] + "&page=" + page)

    } else if (values[2] == "http://localhost:8080/categories") {
        window.location.replace("/categories?perpage=" + values[0] + "&page=" + page)

    } else {
        window.location.replace("/")
    }
}

function perPage(quantity) {
    values = valueFromUrl()

    perpage = quantity

    if (values[2] == "http://localhost:8080/posts/content") {
        window.location.replace("/posts/content?post=" + values[3] + "&perpage=" + perpage + "&page=" + values[1])

    } else if (values[2] == "http://localhost:8080/posts") {
        window.location.replace("/posts?category=" + values[3] + "&perpage=" + perpage + "&page=" + values[1])

    } else if (values[2] == "http://localhost:8080/categories") {
        window.location.replace("/categories?perpage=" + perpage + "&page=" + values[1])

    } else {
        window.location.replace("/")
    }
}

let theSelect = document.querySelector('#perpage')

theSelect.addEventListener( 'change', () => {
    choice = theSelect.selectedIndex;
    value = theSelect.options[choice].value
    perPage(value)

}, true)