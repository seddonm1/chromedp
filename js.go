package chromedp

const (
	// textJS is a javascript snippet that returns the concatenated textContent
	// of all visible (ie, offsetParent !== null) children.
	textJS = `(function(a) {
		var s = '';
		for (var i = 0; i < a.length; i++) {
			if (a[i].offsetParent !== null) {
				s += a[i].textContent;
			}
		}
		return s;
	})(%s.children)`

	// blurJS is a javscript snippet that blurs the specified element.
	blurJS = `(function(a) {
		a.blur();
		return true;
	})(%s)`

	// scrollIntoViewJS is a javascript snippet that scrolls the specified node
	// into the window's viewport (if needed), returning the actual window x/y
	// after execution.
	scrollIntoViewJS = `(function(a) {
		a.scrollIntoViewIfNeeded(true);
		return [window.scrollX, window.scrollY];
	})(%s)`

	// submitJS is a javascript snippet that will call the containing form's
	// submit function, returning true or false if the call was successful.
	submitJS = `(function(a) {
		if (a.nodeName === 'FORM') {
			a.submit();
			return true;
		} else if (a.form !== null) {
			aform.submit();
			return true;
		}
		return false;
	})(%s)`

	// resetJS is a javascript snippet that will call the containing form's
	// reset function, returning true or false if the call was successful.
	resetJS = `(function(a) {
		if (a.nodeName === 'FORM') {
			a.reset();
			return true;
		} else if (a.form !== null) {
			a.form.reset();
			return true;
		}
		return false;
	})(%s)`

	// attributeJS is a javascript snippet that returns the attribute of a specified
	// node.
	attributeJS = `(function(a, n) {
		return a[n];
	})(%s, %s)`

	// setAttributeJS is a javascript snippet that sets the value of the specified
	// node, and returns the value.
	setAttributeJS = `(function(a, n, v) {
		return a[n] = v;
	})(%s, %s, %s)`

	// visibleJS is a javascript snippet that returns true or false depending
	// on if the specified node's offsetParent is not null.
	visibleJS = `(function(a) {
		return a.offsetParent !== null;
	})(%s)`
)
