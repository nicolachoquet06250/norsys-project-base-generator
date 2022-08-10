const buildTree = (tree, basePath, pathSeparator, isRoot) => {
    let item, mainList

    if (isRoot) {
        const handleRootClick = () => buildTreeInBody(basePath, pathSeparator, tree);

        mainList = document.createElement('ul')

        item = document.createElement('li')
        item.classList.add('opened')
        item.style.cursor = 'pointer'
        item.style.paddingBottom = '5px';

        const folderIcon = document.createElement('i');
        folderIcon.classList.add('fa-solid', 'fa-house')
        folderIcon.addEventListener('click', handleRootClick)
        item.appendChild(folderIcon)

        const text = document.createElement('span')
        text.innerHTML = 'Home'
        text.addEventListener('click', handleRootClick)
        item.appendChild(text)

        mainList.appendChild(item)
    }

    const list = document.createElement('ul')

    for (const dir of (tree ?? [])) {
        const item = document.createElement('li')
        item.setAttribute('data-dir', (basePath + pathSeparator + dir).replaceAll(/[:\/\\ ]/g, '-'))
        item.setAttribute('data-path', basePath + pathSeparator + dir)
        item.classList.add('closed', 'tree-item')

        const folderIcon = document.createElement('i');
        folderIcon.classList.add('fa-solid', 'fa-folder')
        folderIcon.addEventListener('click', onItemClicked)
        item.appendChild(folderIcon)

        const text = document.createElement('span')
        text.innerHTML = dir
        text.addEventListener('click', onItemClicked)
        item.appendChild(text)

        const input = document.createElement('input')
        input.setAttribute('type', 'hidden')
        input.value = basePath + pathSeparator + dir
        item.appendChild(input)

        const subList = document.createElement('ul')
        item.appendChild(subList)

        list.appendChild(item)
    }

    if (isRoot) {
        item.appendChild(list)

        return mainList
    }

    return Array.from(list.children)
};

const buildTreeInBody = (basePath, pathSeparator, tree) => {
    const windowBody = document.querySelector('.element-list');

    windowBody.innerHTML = '';

    for (const folder of (tree ?? [])) {
        const f = document.createElement('div')

        const icon = document.createElement('i')
        icon.classList.add('fa-solid', 'fa-folder')
        f.appendChild(icon)

        const text = document.createElement('span');
        text.innerHTML = folder
        f.appendChild(text)

        const input = document.createElement('input')
        input.setAttribute('type', 'hidden');
        input.value = basePath + pathSeparator + folder;
        f.appendChild(input)

        f.addEventListener('click', () => {
            document.querySelector('.element-list .selected')?.classList.remove('selected')
            f.classList.add('selected');

            document.querySelector('#folder').value = basePath + pathSeparator + folder;
        });

        windowBody.appendChild(f)
    }
}

const onItemClicked = e => {
    const element = e.target.parentElement;
    const list = element.querySelector('ul')

    sendMessage({
        channel: 'OpenFolder',
        data: {
            folder: element.getAttribute('data-path')
        }
    })

    if (element.classList.contains('opened')) {
        element.classList.remove('opened');
        element.classList.add('closed');
    } else {
        element.classList.remove('closed');
        element.classList.add('opened');
    }
};

document.addEventListener('astilectron-ready', () => {
    /**
     * @param {string} message.channel
     * @param {Record<string, string>|Record<string, string>[]} message.data
     * @returns {string}
     */
    const handleNewMessage = message => {
        switch (message.channel) {
            case 'GetTree':
                const basePath = message.data.basePath;
                const pathSeparator = message.data.pathSeparator;
                const isRoot = message.data.isHome;
                const tree = message.data.tree;

                const element = document.querySelector(`li[data-dir="${basePath.replaceAll(/[:\/\\ ]/g, '-')}"]`);
                const list = element?.querySelector('ul') ?? { children : [] }

                if (list.children.length === 0) {
                    let root = document.querySelector('.left-menu > nav')
                        .querySelector(`[data-dir="${basePath.replaceAll(/[:\/\\ ]/g, '-')}"]`)
                        ?.querySelector('ul');

                    if (isRoot) {
                        root = document.querySelector('.left-menu > nav');
                    }

                    const builtTree = buildTree(tree, basePath, pathSeparator, isRoot)

                    if (isRoot) {
                        root.appendChild(builtTree)
                    } else {
                        for (const element of builtTree) root.appendChild(element)
                    }
                }

                buildTreeInBody(basePath, pathSeparator, tree)

                return tree;
        }
    };
    astilectron.onMessage(handleNewMessage);

    document.querySelector('.element-list').addEventListener('click', e => {
        if (e.target.classList.contains('element-list')) {
            document.querySelector('.element-list .selected')?.classList.remove('selected')
        }
    })

    setTimeout(() => {
        sendMessage({ channel: `OpenFolder` })
    }, 1000);

    const handleClick = () => {
        const folderElement = document.querySelector('#folder')

        if (folderElement.value !== '') {
            sendMessage({
                channel: 'ChooseFolder',
                data: {
                    path: folderElement.value
                }
            })
        }
    };
    document.querySelector('#open').addEventListener('click', handleClick)
});