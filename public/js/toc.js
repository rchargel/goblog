(function($) {
	$.fn.searchBox = function() {
		this.each(function() {
			var textbox = $(this),
			    bodyArticle = $('#body-content'),
				searchArticle = $('#search-section'),
				url = '/search';

			textbox.bind('keyup', function() {
				if (textbox.val().length > 2) {
					searchArticle.load(url, {terms: textbox.val()}, function() {
						bodyArticle.hide();
						searchArticle.show();
					});
				} else {
					bodyArticle.show();
					searchArticle.hide();
				}
			});
			textbox.triggerHandler('keyup');
		});
		return this;
	};
	$.fn.toc = function() {
		this.each(function() {
			var ul = $(this), headers = ul.nextAll(":header"), lastLevel = -1, toc = {
				id : 'root',
				name : 'root',
				level : 0,
				parent : null,
				children : []
			}, parent = toc;

			var push = function(parent, el, level) {
				parent.children.push({
					id : el.attr('id'),
					name : el.text(),
					level : level,
					parent : parent,
					children : []
				})
			};

			var addItems = function(ul, children) {
				var i = 0; li = null, child = null, sul = null;
				for (; i < children.length; i++) {
					child = children[i];
					li = $('<li></li>');
					li.append('<a href="#' + child.id + '">'+child.name+'</a>');
					ul.append(li);

					if (child.children.length > 0) {
						sul = $('<ul></ul>');
						li.append(sul);
						addItems(sul, child.children);
					}
				}
			}

			headers.each(function() {
				if (this) {
					var header = this, level = parseInt((header.localName || header.nodeName).substring(1));

					if ($(header).attr('id') !== null) {
						if (lastLevel === -1 || lastLevel === level) {
							push(parent, $(header), level);
						} else if (lastLevel < level) {
							parent = parent.children[parent.children.length - 1];
							push(parent, $(header), level);
						} else {
							parent = parent.parent;
							while (parent.level >= level) {
								parent = parent.parent;
							}

							push(parent, $(header), level);
						}
						lastLevel = level;
					}
				}
			});

			addItems(ul, toc.children);
		})
		return this;
	};

	$(document).ready(function() {
		var toc = $('.table-of-contents');
		if (toc.size() > 0) {
			toc.toc();
		}
		$('.searchbox').searchBox();
		$('figure#menu').click(function() {
			var menu = $('aside.main > nav');
			if (menu.is(':visible')) {
				menu.removeAttr('style');
			} else {
				menu.attr('style', 'display:block;');
			}
		})
	});
})(jQuery);
