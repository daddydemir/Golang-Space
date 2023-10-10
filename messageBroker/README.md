## Kafka Message Broker
----

![image](https://github.com/daddydemir/Golang-Space/assets/42716195/084e3f93-7d90-49d6-9539-86ff8228951f)


----
Topic ler icerisinde partitionlar olusturuluyor (kac tane partition olacagi sayi olarak belirtiliyor. ornek olarak 3) Daha sonra mesajlar produce ediliyor. Consumer tarafinda ise bir consumer group olusturuluyor (groupID ye atama yapilarak bu adim tamamlanmis 
oluyor) ve ilgili topicden gelen mesajlar consume ediliyor. 
Partition sayisi 3 oldugu icin maksimum consumer sayisi da 3 oluyor. eger iki tane yada bir tane consumer ayaga kaldirilirsa butum mesajlari o consumer/consumerlara iletiliyor. 



