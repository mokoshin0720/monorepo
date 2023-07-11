import React from 'react';
import { StyleSheet, TouchableOpacity, ViewStyle } from 'react-native';
import { MaterialIcons } from '@expo/vector-icons';
import { baseColor, mainColor } from 'src/utils/color';

type CircleButtonProps = {
    name: keyof typeof MaterialIcons.glyphMap;
    onPress: () => void;
    style?: ViewStyle;
};

export const CircleButton: React.FC<CircleButtonProps> = ({
    name,
    onPress,
    style,
}) => {
    return (
        <TouchableOpacity
            style={[styles.circleButton, style]}
            onPress={onPress}
        >
            <MaterialIcons name={name} size={32} color="white" />
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    circleButton: {
        backgroundColor: mainColor,
        width: 64,
        height: 64,
        borderRadius: 32,
        justifyContent: 'center',
        alignItems: 'center',
        position: 'absolute',
        right: 40,
        bottom: 40,
        shadowColor: '#000000',
        shadowOffset: { width: 0, height: 8 },
        shadowOpacity: 0.25,
        shadowRadius: 8,
        elevation: 8,
    },
    circleButtonLabel: {
        color: baseColor,
        fontSize: 40,
        lineHeight: 40,
    },
});
