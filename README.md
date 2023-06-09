# lineman-deliver-orders

input ข้อมูลตามไฟล์ .txt</br>
output เป็นเป็นทางที่สั้นที่สุดที่จะไปถึงจุดหมาย

start point จุดเริ่มโปรแกรม</br>
end point จุดสิ้นสุดโปรแกรม

ใช้ dijkstra algorithm ในการหาเส้นทางที่สั้นที่สุด</br>
การใช้ dijkstra algorithm ต้องมีการเก็บข้อมูลในรูปแบบ stack ตามการอ้างอิงจาก blog ดังนี้
https://medium.com/@weeklycpproblems/tips-tricks-for-toi-ep-1-dijkstra-shortest-path-de30ed55c9a1

Heap เป็นโครงสร้างข้อมูลที่นํามาสร้างแถวคอยลําดับความสําคัญ(priority queue)
http://www.ecpe.nu.ac.th/panomkhawn/datastandal/pdf/lecture09.pdf