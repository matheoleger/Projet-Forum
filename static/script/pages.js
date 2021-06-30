function valueFromUrl() {
	url = window.location.href
    begin_url = url.match(/^.*(?=(\?))/gm).toString()

    post = url.match(/[?&]+post=?([^&]*)?/gi).toString()
    postv = post.match(/\d/gi).toString()
    postvalue = postv.replace(/,/g, '')

    category = url.match(/[?&]+category=?([^&]*)?/gi)
    if (category != null ) {
        categoryv = category.toString()
        categoryvalue = category.replace(/category=/g, '')
    }

    perpage = url.match(/[?&]+perpage=?([^&]*)?/gi).toString()
    perpagev = perpage.match(/\d/gi).toString()
    perpagevalue = perpagev.replace(/,/g, '')

    page = url.match(/[?&]+page=?([^&]*)?/gi).toString()
    pagev = page.match(/\d/gi).toString()
    pagevalue = pagev.replace(/,/g, '')

    console.log("postvalue=" + postvalue + "perpagevalue=" + perpagevalue + "pagevalue=" + pagevalue )
    return [postvalue, perpagevalue, pagevalue, begin_url]
}

function nextPage() {
    values = valueFromUrl()

    page = parseInt(values[2]) + 1

    console.log("next page" + values)

    if (values[3] = "http://localhost:8080/posts/content") {
        window.location.replace("/posts/content?post=" + values[0] + "&perpage=" + values[1] + "&page=" + page)
    
    } else if (values[3] = "http://localhost:8080/posts") {
        window.location.replace("/posts?category=Autres&perpage=" + values[1] + "&page=" + page)

    } else if (values[3] = "http://localhost:8080/categories") {
        window.location.replace("/categories?perpage=" + values[1] + "&page=" + page)

    } else {
        window.location.replace("/404")
    }

    // le good regex = /[?&]+page=?([^&]*)?/gi
}

function previousPage() {
    values = valueFromUrl()

    page = values[2]
    if (page > 0) {
        page--
    }

    window.location.replace("/posts/content?post=" + values[0] + "&perpage=" + values[1] + "&page=" + page)
}

function perPage(quantity) {
    values = valueFromUrl()

    perpage = quantity

    window.location.replace("/posts/content?post=" + values[0] + "&perpage=" + quantity + "&page=" + values[2])
}

let theSelect = document.querySelector('#perpage')

theSelect.addEventListener( 'change', () => { // onchange="perPage()
    // values = valueFromUrl()
    choice = theSelect.selectedIndex;
    value = theSelect.options[choice].value
    perPage(value)

}, true)