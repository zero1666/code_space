/*
	读取配置文件，
    支持注释，注释前导符为双斜杠（跟c++的单行注释一样)://
	支持[section]章节
	如:
	[gmcc] # gmcc section
	[unicom] # unicom section
*/

#if !defined _SERVER_CONFIG_H_
#define _SERVER_CONFIG_H_

#include <fstream>
#include <map>
#include <string>
#include <string.h>

using namespace std;

class CServerConf {
public:
	CServerConf();
	CServerConf(const char* szFileName);
	virtual ~CServerConf();


public:
	const string& operator[](const char* szName);
	const string& operator[](const string& strName);
	const string& operator()(const char* szSection, const char* szName);

	int ParseFile(const char* szConfigFile);
	static int StrimString(char* szLine);

private:

	int ParseFile();

private:
	ifstream m_ConfigFile;
	map<string, string> m_ConfigMap;
};


#endif

