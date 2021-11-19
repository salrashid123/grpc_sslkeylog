package com.test;

import com.google.cloud.pubsub.v1.TopicAdminClient;
import com.google.cloud.pubsub.v1.TopicAdminSettings;
import com.google.pubsub.v1.ProjectName;
import com.google.pubsub.v1.Topic;

public class TestApp {
	public static void main(String[] args) {
		TestApp tc = new TestApp();
	}

	private static String projectID = "mineral-minutia-820";

	public TestApp() {
		try {

			// ConsoleHandler consoleHandler = new ConsoleHandler();
			// consoleHandler.setLevel(Level.ALL);
			// consoleHandler.setFormatter(new SimpleFormatter());

			// FileHandler fh = new FileHandler("/tmp/grpc.log");

			// Logger logger = Logger.getLogger("com.google.api.client");
			// logger.setLevel(Level.ALL);
			// logger.addHandler(consoleHandler);

			// Logger lh = Logger.getLogger("httpclient.wire.header");
			// lh.setLevel(Level.ALL);
			// lh.addHandler(consoleHandler);

			// Logger lc = Logger.getLogger("httpclient.wire.content");
			// lc.setLevel(Level.ALL);
			// lc.addHandler(consoleHandler);

			// Logger lc2 = Logger.getLogger("io.grpc");
			// lc2.setLevel(Level.FINE);
			// lc2.addHandler(consoleHandler);
			// lc2.addHandler(fh);

			TopicAdminSettings topicAdminSettings = TopicAdminSettings.newBuilder().build();
			TopicAdminClient topicAdminClient = TopicAdminClient.create(topicAdminSettings);
			ProjectName project = ProjectName.of(projectID);
			for (Topic element : topicAdminClient.listTopics(project).iterateAll())
				System.out.println(element.getName());

		} catch (Exception ex) {
			System.out.println("Error:  " + ex.getMessage());
		}
	}

}
