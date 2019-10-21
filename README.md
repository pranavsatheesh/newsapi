Clone the repository using : git clone https://github.com/pranavsatheesh/newsapi.git

After cloning the project. 

In conf -> app.conf
            
            Edit app.conf
                 httpport = 1024 to available port number.
                 Db details as per your local configurations
                    dbDriver=mysql
                    dbUser=root
                    dbPass=root
                    dbName=news
Run build file ./news
It automatically creates database. (Make sure parameters in app.conf is valid. Database structure is also available in prokect folder news_db.sql)

To get the view page either use 'http://localhost:1024/index' or use index.html.
