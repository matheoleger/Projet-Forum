function valueFromUrl() {
	url = window.location.href
    post = url.match(/[?&]+post=?([^&]*)?/gi).toString()
    postv = post.match(/\d/gi).toString()
    postvalue = postv.replace(/,/g, '')
    perpage = url.match(/[?&]+perpage=?([^&]*)?/gi).toString()
    perpagev = perpage.match(/\d/gi).toString()
    perpagevalue = perpagev.replace(/,/g, '')
    page = url.match(/[?&]+page=?([^&]*)?/gi).toString()
    pagev = page.match(/\d/gi).toString()
    pagevalue = pagev.replace(/,/g, '')

    console.log("postvalue=" + postvalue + "perpagevalue=" + perpagevalue + "pagevalue=" + pagevalue )
    return [postvalue, perpagevalue, pagevalue]
}

function nextPage() {
    values = valueFromUrl()

    page = parseInt(values[2]) + 1

    console.log("next page" + values)
    window.location.replace("/posts/content?post=" + values[0] + "&perpage=" + values[1] + "&page=" + page)


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