import {createApp} from "vue"

const database = createApp({}).config.globalProperties

const name = "uncle_first"

/**
 * 初始化数据库
 */
const initDatabase = () => {
	if (!sessionStorage.getItem(name)) {
		sessionStorage.setItem(name, JSON.stringify({}))
	}
}

/**
 * 获取数据库
 */
const getDatabase = () => {
	initDatabase()
	return JSON.parse(sessionStorage.getItem(name))
}

/**
 * 保存数据库
 */
const saveDatabase = (data) => {
	sessionStorage.setItem(name, JSON.stringify(data))
}

/**
 * 获取数据库中的值
 * @param key 数据库中的键
 * @param defaultValue 默认值
 */
database.get = (key, defaultValue = null) => {
	const db = getDatabase()
	if (db.hasOwnProperty(key)) {
		return db[key]
	}
	// 不存在就调用add创建
	return database.add(key, defaultValue)
}

/**
 * 添加/更新数据库中的值
 * @param key 数据库中的键
 * @param value 数据库中的值
 */
database.add = (key, value) => {
	const db = getDatabase()
	db[key] = value
	saveDatabase(db)
	return value
}

/**
 * 删除数据库中的值
 * @param key 数据库中的键
 */
database.delete = (key) => {
	const db = getDatabase()
	if (db.hasOwnProperty(key)) {
		delete db[key]
		saveDatabase(db)
		return true
	}
	return false
}

// 初始化数据库
initDatabase()

export default database