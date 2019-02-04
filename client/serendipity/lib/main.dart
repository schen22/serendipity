import 'package:flutter/material.dart';

import 'chat_screen.dart';

/// SerendipityApp is Flutter application
class SerendipityApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Serendipity",
      home: ChatScreen(),
    );
  }
}

/// main is entry point of Flutter application
void main() {
  runApp(SerendipityApp());
}
