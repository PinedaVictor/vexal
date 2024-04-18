import user.User;
import user.HairColor;
import data.TestingData;
import org.apache.spark
import akka.actor.typed.scaladsl.Behaviors
import akka.actor.typed.scaladsl.LoggerOps
import akka.actor.typed.{ActorRef, ActorSystem, Behavior}

object testAPI {
  println("Is this run");
  println(Behavior);
}

@main def main = {
  println("In the main funciton");

  val oneUser = TestingData.generateUser();
  println(oneUser.getFullName());
  val multipleUsers = TestingData.generateUserData(12);
  println("Number of users: " + multipleUsers.length);

  // for i <- multipleUsers yield println(s"${i.getAge()} ${i.getFullName()} ${i.getUserID()}");

}